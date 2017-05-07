package gopatrol

import (
	"time"
)

type Endpoint struct {
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
	Attempts int `json:"attempts,omitempty"`
}
