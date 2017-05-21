package gopatrol

type CheckersService interface {
	InsertChecker(checker interface{}) error
	GetAllCheckers(query map[string]interface{}) ([]interface{}, error)
	GetCheckerBySlug(slug string) (interface{}, error)
	CountCheckers(query map[string]interface{}) (int, error)
	DeleteCheckerBySlug(slug string) error
	UpdateCheckerBySlug(slug string, updateData interface{}) error
}

type EventService interface {
	InsertEvent(event Event) error
	GetAllEvents(query map[string]interface{}) ([]Event, error)
	CountEvents(query map[string]interface{}) (int, error)
}

type CacheService interface {
	// serialize to bytes
	Set(key string, value []byte) error
	Get(key string) ([]byte, error)
}

type User struct {
	Name         string `json:"name" bson:"name" valid:"required,email"`
	Email        string `json:"email" bson:"email" valid:"required,email"`
	HashPassword string `json:"-" bson:"hashpassword"`
	Role         string `json:"role" bson:"role"`
}

type UsersService interface {
	Register(name, email, role, plainPassword string) error
	// UpdateUser(email string, updateData interface{}) error
	Login(email, plainPassword string) (User, error)
}
