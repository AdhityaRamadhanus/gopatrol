package gopatrol

import "time"

type Event struct {
	URL       string    `json:"url"`
	Slug      string    `json:"slug"`
	Type      string    `json:"type"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}
