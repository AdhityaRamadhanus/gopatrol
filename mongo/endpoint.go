package mongo

import (
	"github.com/AdhityaRamadhanus/gopatrol/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CheckersService struct {
	session  *mgo.Session
	CollName string
}

func NewCheckersService(session *mgo.Session, collName string) *CheckersService {
	// Ensure Index
	slugIdx := mgo.Index{
		Key:        []string{"slug"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	urlIdx := mgo.Index{
		Key:        []string{"url"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	session.DB(config.DatabaseName).C(collName).EnsureIndex(slugIdx)
	session.DB(config.DatabaseName).C(collName).EnsureIndex(urlIdx)
	return &CheckersService{
		session:  session,
		CollName: collName,
	}
}

// for more flexible use
func (p *CheckersService) CopySession() *mgo.Session {
	return p.session.Copy()
}

func (p *CheckersService) InsertChecker(checker interface{}) error {
	copySession := p.session.Copy()
	defer copySession.Close()
	CheckersColl := copySession.DB(config.DatabaseName).C(p.CollName)
	return CheckersColl.Insert(checker)
}

func (p *CheckersService) UpdateCheckerBySlug(slug string, UpdateData interface{}) error {
	copySession := p.session.Copy()
	defer copySession.Close()
	CheckersColl := copySession.DB(config.DatabaseName).C(p.CollName)
	return CheckersColl.Update(bson.M{"slug": slug}, UpdateData)
}

func (p *CheckersService) GetAllCheckers(q map[string]interface{}) ([]interface{}, error) {
	copySession := p.session.Copy()
	defer copySession.Close()

	CheckersColl := copySession.DB(config.DatabaseName).C(p.CollName)
	endpoints := []interface{}{}

	MongoQuery := CheckersColl.Find(q["query"])
	if ok, val := q["pagination"].(bool); ok && val {
		MongoQuery.
			Skip(q["page"].(int) * q["limit"].(int)).
			Limit(q["limit"].(int))
	}

	if err := MongoQuery.
		All(&endpoints); err != nil {
		return nil, err
	}
	return endpoints, nil
}

func (p *CheckersService) GetCheckerBySlug(slug string) (interface{}, error) {
	copySession := p.session.Copy()
	defer copySession.Close()

	CheckersColl := copySession.DB(config.DatabaseName).C(p.CollName)
	var endpoint interface{}
	if err := CheckersColl.
		Find(bson.M{
			"slug": slug,
		}).
		One(&endpoint); err != nil {
		return nil, err
	}
	return endpoint, nil
}

func (p *CheckersService) DeleteCheckerBySlug(slug string) error {
	copySession := p.session.Copy()
	defer copySession.Close()

	CheckersColl := copySession.DB(config.DatabaseName).C(p.CollName)
	return CheckersColl.Remove(bson.M{"slug": slug})
}
