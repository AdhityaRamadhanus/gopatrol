package daemon

import (
	"sync"
	"time"

	checkup "github.com/AdhityaRamadhanus/gopatrol"
	"github.com/AdhityaRamadhanus/gopatrol/mongo"
	log "github.com/Sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//ServiceHandler is a grpc server and checkup server
type Daemon struct {
	*checkup.Checkup
	globalLock      sync.RWMutex
	CheckInterval   time.Duration
	EndpointService checkup.EndpointService
}

//NewServiceHandler create new ServiceHandler from a configfile (checkup.json)
func NewDaemon(mongoSession *mgo.Session) (*Daemon, error) {
	daemon := &Daemon{
		Checkup: &checkup.Checkup{},
	}
	daemon.Storage = mongo.NewLoggingService(mongoSession)
	daemon.EndpointService = mongo.NewEndpointService(mongoSession)
	endpoints, err := daemon.EndpointService.GetAllEndpoints(map[string]interface{}{
		"query": bson.M{},
	})
	if err != nil {
		log.WithError(err).Error("Initialize daemon failed")
		return nil, err
	}
	// Load the checkers to memory
	for _, endpoint := range endpoints {
		typeEndpoint := endpoint.(bson.M)["type"]
		switch typeEndpoint {
		case "tcp":
			var tcpChecker checkup.TCPChecker
			bsonBytes, err := bson.Marshal(endpoint)
			log.WithError(err).Error(ErrFailedToMarshalBSON)
			if err := bson.Unmarshal(bsonBytes, &tcpChecker); err != nil {
				log.WithError(err).Error(ErrFailedToUnmarshalBSON)
			}
			daemon.Checkers = append(daemon.Checkers, tcpChecker)
		case "http":
			var httpChecker checkup.HTTPChecker
			bsonBytes, err := bson.Marshal(endpoint)
			log.WithError(err).Error(ErrFailedToMarshalBSON)
			if err := bson.Unmarshal(bsonBytes, &httpChecker); err != nil {
				log.WithError(err).Error(ErrFailedToUnmarshalBSON)
			}
			daemon.Checkers = append(daemon.Checkers, httpChecker)
		case "dns":
			var dnsChecker checkup.DNSChecker
			bsonBytes, err := bson.Marshal(endpoint)
			log.WithError(err).Error(ErrFailedToMarshalBSON)
			if err := bson.Unmarshal(bsonBytes, &dnsChecker); err != nil {
				log.WithError(err).Error(ErrFailedToUnmarshalBSON)
			}
			daemon.Checkers = append(daemon.Checkers, dnsChecker)
		}
	}
	return daemon, nil
}

//Run the main loop for checkup to check endpoints and should be run as a goroutine
func (d *Daemon) Run() {
	for {
		timer := time.After(d.CheckInterval)
		select {
		case <-timer: //hardcoded for now
			// Obtain Lock, makesure no function updating the Checkers
			d.globalLock.RLock()
			results, err := d.Check()
			if err != nil {
				log.WithError(err).Error(ErrFailedToDoCheck)
			} else {
				if err := d.Storage.Store(results); err != nil {
					log.WithError(err).Error(ErrFailedToStoreResult)
				}
				if err := d.CheckEventsAndSync(results); err != nil {
					log.WithError(err).Error(ErrFailedToStoreResult)
				}
			}
			d.globalLock.RUnlock()
		}
	}
}

func (d *Daemon) CheckEventsAndSync(results []checkup.Result) error {
	var events []checkup.Event
	var err error
	for _, result := range results {
		var c checkup.Checker
		for _, checker := range d.Checkers {
			if checker.GetSlug() == result.Slug {
				c = checker
			}
		}

		switch {
		case result.Down:
			if c.GetLastStatus() == "healthy" {
				events = append(events, checkup.Event{
					Message:   c.GetURL() + " is down",
					Type:      "down",
					URL:       c.GetURL(),
					Slug:      c.GetSlug(),
					Timestamp: result.Timestamp,
				})
			}
		case result.Healthy:
			if c.GetLastStatus() == "down" {
				events = append(events, checkup.Event{
					Message:   c.GetURL() + " is up",
					Type:      "up",
					URL:       c.GetURL(),
					Slug:      c.GetSlug(),
					Timestamp: result.Timestamp,
				})
			}
		}

		// Update Status in MONGODB to sync with the memoery

		var updateData = bson.M{
			"$set": bson.M{
				"last_checked": result.Timestamp,
				"last_status":  result.Status(),
			},
		}
		err = d.EndpointService.UpdateEndpointBySlug(result.Slug, updateData)
	}

	return err
}
