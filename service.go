package gopatrol

type CheckersService interface {
	InsertChecker(checker interface{}) error
	GetAllCheckers(query map[string]interface{}) ([]interface{}, error)
	GetCheckerBySlug(slug string) (interface{}, error)
	DeleteCheckerBySlug(slug string) error
	UpdateCheckerBySlug(slug string, updateData interface{}) error
}

type EventService interface {
	InsertEvent(event Event) error
	GetAllEvents(query map[string]interface{}) ([]Event, error)
}

type CacheService interface {
	// serialize to bytes
	Set(key string, value []byte) error
	Get(key string) ([]byte, error)
}
