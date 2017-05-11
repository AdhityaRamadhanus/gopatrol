package gopatrol

import (
	"fmt"
	"net"
	"time"

	"github.com/miekg/dns"
)

// DNSChecker implements a Checker for TCP endpoints.
type DNSChecker struct {
	Slug string `json:"slug" valid:"required"`
	// Name is the name of the endpoint.
	Name string `json:"name,omitempty" valid:"required"`
	// URL is the URL of the endpoint.
	URL  string `json:"url,omitempty" valid:"required"`
	Type string `json:"type,omitempty" valid:"required"`
	// ThresholdRTT is the maximum round trip time to
	// allow for a healthy endpoint. If non-zero and a
	// request takes longer than ThresholdRTT, the
	// endpoint will be considered unhealthy. Note that
	// this duration includes any in-between network
	// latency.
	ThresholdRTT time.Duration `json:"threshold_rtt,omitempty"`
	// Attempts is how many requests the client will
	// make to the endpoint in a single check.
	Attempts int    `json:"attempts,omitempty"` // This is the fqdn of the target server to query the DNS server for.
	Host     string `json:"hostname_fqdn,omitempty"`
	// Timeout is the maximum time to wait for a
	// TCP connection to be established.
	Timeout     time.Duration `json:"timeout,omitempty"`
	LastChecked time.Time     `json:"last_checked"`
	LastStatus  string        `json:"last_status"`
}

func (c DNSChecker) GetName() string {
	return c.Name
}

func (c DNSChecker) GetURL() string {
	return c.URL
}

func (c DNSChecker) GetSlug() string {
	return c.Slug
}

func (c DNSChecker) GetLastChecked() time.Time {
	return c.LastChecked
}

func (c DNSChecker) GetLastStatus() string {
	return c.LastStatus
}

// Check performs checks using c according to its configuration.
// An error is only returned if there is a configuration error.
func (c DNSChecker) Check() (Result, error) {
	if c.Attempts < 1 {
		c.Attempts = 1
	}

	result := Result{Name: c.Name, URL: c.URL, Timestamp: time.Now().UTC(), Slug: c.Slug}
	result.Times = c.doChecks()

	return c.conclude(result), nil
}

// doChecks executes and returns each attempt.
func (c DNSChecker) doChecks() Attempts {
	var err error
	var conn net.Conn

	timeout := c.Timeout
	if timeout == 0 {
		timeout = 1 * time.Second
	}

	checks := make(Attempts, c.Attempts)
	for i := 0; i < c.Attempts; i++ {
		start := time.Now()

		if c.Host != "" {
			hostname := c.Host
			m1 := new(dns.Msg)
			m1.Id = dns.Id()
			m1.RecursionDesired = true
			m1.Question = make([]dns.Question, 1)
			m1.Question[0] = dns.Question{hostname, dns.TypeA, dns.ClassINET}
			d := new(dns.Client)
			if err != nil {
				checks[i].Error = err.Error()
				continue
			}
			_, _, err := d.Exchange(m1, c.URL)
			if err != nil {
				checks[i].Error = err.Error()
				continue
			}
		}
		if conn, err = net.DialTimeout("tcp", c.URL, c.Timeout); err == nil {
			conn.Close()
		}

		checks[i].RTT = time.Since(start)
		if err != nil {
			checks[i].Error = err.Error()
			continue
		}
	}
	return checks
}

// conclude takes the data in result from the attempts and
// computes remaining values needed to fill out the result.
// It detects degraded (high-latency) responses and makes
// the conclusion about the result's status.
func (c DNSChecker) conclude(result Result) Result {
	result.ThresholdRTT = c.ThresholdRTT

	// Check errors (down)
	for i := range result.Times {
		if result.Times[i].Error != "" {
			result.Down = true
			return result
		}
	}

	// Check round trip time (degraded)
	if c.ThresholdRTT > 0 {
		stats := result.ComputeStats()
		if stats.Median > c.ThresholdRTT {
			result.Notice = fmt.Sprintf("median round trip time exceeded threshold (%s)", c.ThresholdRTT)
			result.Degraded = true
			return result
		}
	}

	result.Healthy = true
	return result
}
