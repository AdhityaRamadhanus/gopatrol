// Package checkup provides means for checking and reporting the
// status and performance of various endpoints in a distributed,
// lock-free, self-hosted fashion.
package gopatrol

import (
	"fmt"
	"sync"
	"time"
)

// Checker can create a Result.
type Checker interface {
	Check() (Result, error)
}

// Notifier can notify ops or sysadmins of
// potential problems. A Notifier should keep
// state to avoid sending repeated notices
// more often than the admin would like.
type Notifier interface {
	Notify(Result) error
}

// DefaultConcurrentChecks is how many checks,
// at most, to perform concurrently.
var DefaultConcurrentChecks = 5

// Checkup performs a routine checkup on endpoints or
// services.
type Checkup struct {
	// Checkers is the list of Checkers to use with
	// which to perform checks.
	Checkers []Checker `json:"checkers,omitempty"`

	// ConcurrentChecks is how many checks, at most, to
	// perform concurrently. Default is
	// DefaultConcurrentChecks.
	ConcurrentChecks int `json:"concurrent_checks,omitempty"`

	// Timestamp is the timestamp to force for all checks.
	// Useful if wanting to perform distributed check
	// "at the same time" even if they might actually
	// be a few milliseconds or seconds apart.
	Timestamp time.Time `json:"timestamp,omitempty"`

	// Notifier is a notifier that will be passed the
	// results after checks from all checkers have
	// completed. Notifier may evaluate and choose to
	// send a notification of potential problems.
	Notifier []Notifier `json:"notifier,omitempty"`
}

// Check performs the health checks. An error is only
// returned in the case of a misconfiguration or if
// any one of the Checkers returns an error.
func (c Checkup) Check() ([]Result, error) {
	if c.ConcurrentChecks == 0 {
		c.ConcurrentChecks = DefaultConcurrentChecks
	}
	if c.ConcurrentChecks < 0 {
		return nil, fmt.Errorf("invalid value for ConcurrentChecks: %d (must be set > 0)",
			c.ConcurrentChecks)
	}

	results := make([]Result, len(c.Checkers))
	errs := make(Errors, len(c.Checkers))
	throttle := make(chan struct{}, c.ConcurrentChecks)
	wg := sync.WaitGroup{}

	for i, checker := range c.Checkers {
		throttle <- struct{}{}
		wg.Add(1)
		go func(i int, checker Checker) {
			results[i], errs[i] = checker.Check()
			<-throttle
			wg.Done()
		}(i, checker)
	}
	wg.Wait()

	if !c.Timestamp.IsZero() {
		for i := range results {
			results[i].Timestamp = c.Timestamp
		}
	}

	if !errs.Empty() {
		return results, errs
	}

	// if c.Notifier != nil {
	// 	err := c.Notifier.Notify(results)
	// 	if err != nil {
	// 		return results, err
	// 	}
	// }

	return results, nil
}
