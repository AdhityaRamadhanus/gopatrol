package mongo

import (
	"github.com/AdhityaRamadhanus/gopatrol"
	"github.com/AdhityaRamadhanus/gopatrol/config"
	"gopkg.in/mgo.v2"
)

type EventService struct {
	session  *mgo.Session
	CollName string
}

func NewEventService(session *mgo.Session, collName string) *EventService {
	EventColl := session.DB(config.DatabaseName).C(collName)

	EventColl.Create(&mgo.CollectionInfo{
		Capped: true,
		// Set Max Size in bytes to 5 GB (just a guess number)
		MaxBytes: 5000 * 1000,
		MaxDocs:  100000,
	})

	// Ensure Index
	EventColl.EnsureIndex(mgo.Index{
		Key:        []string{"slug"},
		Unique:     false,
		DropDups:   false,
		Background: true,
		Sparse:     true,
	})

	return &EventService{
		session:  session,
		CollName: collName,
	}
}

// for more flexible use
func (p *EventService) CopySession() *mgo.Session {
	return p.session.Copy()
}

func (p *EventService) InsertEvent(event gopatrol.Event) error {
	copySession := p.session.Copy()
	defer copySession.Close()
	EventColl := copySession.DB(config.DatabaseName).C(p.CollName)
	return EventColl.Insert(event)
}

func (p *EventService) GetAllEvents(q map[string]interface{}) ([]gopatrol.Event, error) {
	copySession := p.session.Copy()
	defer copySession.Close()

	EventColl := copySession.DB(config.DatabaseName).C(p.CollName)
	events := []gopatrol.Event{}
	MongoQuery := EventColl.Find(q["query"])
	if ok, val := q["pagination"].(bool); ok && val {
		MongoQuery.
			Skip(q["page"].(int) * q["limit"].(int)).
			Limit(q["limit"].(int))
	}

	if err := MongoQuery.All(&events); err != nil {
		return nil, err
	}
	return events, nil
}
