package gopatrol

import (
	"time"
)

type EndpointService interface {
	InsertEndpoint(endpoint interface{}) error
	GetAllEndpoints(query map[string]interface{}) ([]interface{}, error)
	GetEndpointBySlug(slug string) (interface{}, error)
	DeleteEndpointBySlug(slug string) error
	UpdateEndpointBySlug(slug string, updateData interface{}) error
}

type LoggingService interface {
	InsertLog(result Result) error
	GetAllLogs(query map[string]interface{}) ([]Result, error)
}

type CacheService interface {
	SetKey(key string, value interface{}, duration time.Duration) error
	GetKey(key string) ([]byte, error)
}
