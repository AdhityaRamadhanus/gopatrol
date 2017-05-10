package mongo

import (
	"log"

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

func (p *LoggingService) GetAllLogs(query interface{}, page, size int) ([]gopatrol.Result, error) {
	copySession := p.session.Copy()
	defer copySession.Close()

	LogColl := copySession.DB(config.DatabaseName).C(collName)
	logs := []gopatrol.Result{}
	if err := LogColl.
		Find(query).
		Skip(page * size).
		Limit(size).
		All(&logs); err != nil {
		return nil, err
	}
	return logs, nil
}

func (p *LoggingService) Store(results []gopatrol.Result) error {
	for _, result := range results {
		if err := p.InsertLog(result); err != nil {
			log.Println("Error Inserting Log", err)
		}
	}
	return nil
}
