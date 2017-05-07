package gopatrol

import "time"

type EndpointService interface {
	InsertEndpoint(endpoint interface{}) error
	GetAllEndpoints(query interface{}, page int, size int) ([]interface{}, error)
	GetEndpointBySlug(slug string) (interface{}, error)
	DeleteEndpointBySlug(slug string) error
	// UpdateEndpoint(query interface{}, updateData interface{}) (interface{}, error)
}

type CacheService interface {
	SetKey(key string, value interface{}, duration time.Duration) error
	GetKey(key string) ([]byte, error)
}
