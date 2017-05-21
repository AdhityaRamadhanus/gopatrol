package mongo

import (
	"github.com/AdhityaRamadhanus/gopatrol"
	"github.com/AdhityaRamadhanus/gopatrol/config"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UsersService struct {
	session  *mgo.Session
	CollName string
}

func NewUsersService(session *mgo.Session, collName string) *UsersService {
	EventColl := session.DB(config.DatabaseName).C(collName)

	// Ensure Index
	EventColl.EnsureIndex(mgo.Index{
		Key:        []string{"email"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	})

	return &UsersService{
		session:  session,
		CollName: collName,
	}
}

// for more flexible use
func (p *UsersService) CopySession() *mgo.Session {
	return p.session.Copy()
}

func (p *UsersService) Register(name, email, role, plainPassword string) error {
	copySession := p.session.Copy()
	defer copySession.Close()
	UserColl := copySession.DB(config.DatabaseName).C(p.CollName)
	user := gopatrol.User{
		Name:  name,
		Email: email,
		Role:  role,
	}
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	user.HashPassword = string(hashPassword)
	return UserColl.Insert(user)
}

func (p *UsersService) Login(email, plainPassword string) (gopatrol.User, error) {
	copySession := p.session.Copy()
	defer copySession.Close()

	UserColl := copySession.DB(config.DatabaseName).C(p.CollName)
	var user gopatrol.User
	if err := UserColl.Find(bson.M{
		"email": email,
	}).One(&user); err != nil {
		return gopatrol.User{}, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.HashPassword), []byte(plainPassword)); err != nil {
		return gopatrol.User{}, err
	}
	return user, nil
}
