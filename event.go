package gopatrol

import "time"

type Event struct {
	Slug      string    `json:"slug" bson:"slug"`
	Name      string    `json:"name" bson:"name"`
	URL       string    `json:"url" bson:"url"`
	Message   string    `json:"message" bson:"message"`
	Type      string    `json:"type" bson:"type"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
	Reason    string    `json:"reason" bson:"reason"`
}

func NewEvent(result Result) Event {
	event := Event{
		Name:      result.Name,
		Slug:      result.Slug,
		URL:       result.URL,
		Timestamp: result.Timestamp,
		Message:   result.Name + " is " + result.Status(),
	}
	if result.Healthy {
		event.Message = result.Name + " is up and running"
		event.Reason = "OK"
	} else {
		event.Message = result.Name + " is down"
		for _, attempt := range result.Times {
			if attempt.Error != "" {
				event.Reason = attempt.Error
				break
			}
		}
	}
	return event
}
