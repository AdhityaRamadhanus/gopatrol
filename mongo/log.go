package mongo

import (
	"github.com/AdhityaRamadhanus/gopatrol"
	"github.com/AdhityaRamadhanus/gopatrol/config"
	"gopkg.in/mgo.v2"
)

const (
	collName = "Logs"
)

type LoggingService struct {
	session  *mgo.Session
	collName string
}

func NewLoggingService(session *mgo.Session) *LoggingService {
	// Ensure Index
	slugIdx := mgo.Index{
		Key:        []string{"slug"},
		Unique:     false,
		DropDups:   false,
		Background: true,
		Sparse:     true,
	}
	session.DB(config.DatabaseName).C(collName).EnsureIndex(slugIdx)
	return &LoggingService{
		session: session,
	}
}

// for more flexible use
func (p *LoggingService) CopySession() *mgo.Session {
	return p.session.Copy()
}

func (p *LoggingService) InsertLog(result gopatrol.Result) error {
	copySession := p.session.Copy()
	defer copySession.Close()
	EndpointColl := copySession.DB(config.DatabaseName).C(collName)
	return EndpointColl.Insert(result)
}

func (p *LoggingService) GetAllLogs(q map[string]interface{}) ([]gopatrol.Result, error) {
	copySession := p.session.Copy()
	defer copySession.Close()

	LogColl := copySession.DB(config.DatabaseName).C(collName)
	logs := []gopatrol.Result{}
	MongoQuery := LogColl.Find(q["query"])
	if ok, val := q["pagination"].(bool); ok && val {
		MongoQuery.
			Skip(q["page"].(int) * q["limit"].(int)).
			Limit(q["limit"].(int))
	}

	if err := MongoQuery.All(&logs); err != nil {
		return nil, err
	}
	return logs, nil
}
