package mongo

import (
	"github.com/AdhityaRamadhanus/gopatrol/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type GetQuery struct {
	Query      bson.M
	Pagination bool
	Page       int
	Limit      int
}

type EndpointService struct {
	session *mgo.Session
}

func NewEndpointService(session *mgo.Session) *EndpointService {
	// Ensure Index
	slugIdx := mgo.Index{
		Key:        []string{"slug"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	session.DB(config.DatabaseName).C("Endpoint").EnsureIndex(slugIdx)
	return &EndpointService{
		session: session,
	}
}

// for more flexible use
func (p *EndpointService) CopySession() *mgo.Session {
	return p.session.Copy()
}

func (p *EndpointService) InsertEndpoint(endpoint interface{}) error {
	copySession := p.session.Copy()
	defer copySession.Close()
	EndpointColl := copySession.DB(config.DatabaseName).C("Endpoint")
	// endpoint.ID = bson.NewObjectId()
	return EndpointColl.Insert(endpoint)
}

func (p *EndpointService) UpdateEndpointBySlug(slug string, UpdateData interface{}) error {
	copySession := p.session.Copy()
	defer copySession.Close()
	EndpointColl := copySession.DB(config.DatabaseName).C("Endpoint")
	return EndpointColl.Update(bson.M{"slug": slug}, UpdateData)
}

func (p *EndpointService) GetAllEndpoints(q map[string]interface{}) ([]interface{}, error) {
	copySession := p.session.Copy()
	defer copySession.Close()

	EndpointColl := copySession.DB(config.DatabaseName).C("Endpoint")
	endpoints := []interface{}{}

	MongoQuery := EndpointColl.Find(q["query"])
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

func (p *EndpointService) GetEndpointBySlug(slug string) (interface{}, error) {
	copySession := p.session.Copy()
	defer copySession.Close()

	EndpointColl := copySession.DB(config.DatabaseName).C("Endpoint")
	var endpoint interface{}
	if err := EndpointColl.
		Find(bson.M{
			"slug": slug,
		}).
		One(&endpoint); err != nil {
		return nil, err
	}
	return endpoint, nil
}

func (p *EndpointService) DeleteEndpointBySlug(slug string) error {
	copySession := p.session.Copy()
	defer copySession.Close()

	EndpointColl := copySession.DB(config.DatabaseName).C("Endpoint")
	return EndpointColl.Remove(bson.M{"slug": slug})
}
