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
	*Endpoint
	// This is the fqdn of the target server to query the DNS server for.
	Host string `json:"hostname_fqdn,omitempty"`
	// Timeout is the maximum time to wait for a
	// TCP connection to be established.
	Timeout time.Duration `json:"timeout,omitempty"`
}

func (c DNSChecker) GetName() string {
	return c.Name
}

func (c DNSChecker) GetURL() string {
	return c.URL
}

// Check performs checks using c according to its configuration.
// An error is only returned if there is a configuration error.
func (c DNSChecker) Check() (Result, error) {
	if c.Attempts < 1 {
		c.Attempts = 1
	}

	result := Result{Title: c.Name, Endpoint: c.URL, Timestamp: Timestamp()}
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
