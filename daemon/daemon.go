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
func (handler *Daemon) Run() {
	for {
		timer := time.After(handler.CheckInterval)
		select {
		case <-timer: //hardcoded for now
			// Obtain Lock, makesure no function updating the Checkers
			handler.globalLock.RLock()
			results, err := handler.Checkup.Check()
			if err != nil {
				log.WithError(err).Error(ErrFailedToDoCheck)
			} else {
				if err := handler.Checkup.Storage.Store(results); err != nil {
					log.WithError(err).Error(ErrFailedToStoreResult)
				}

				for _, result := range results {
					var updateData = bson.M{
						"$set": bson.M{
							"last_checked": result.Timestamp,
							"last_status":  result.Status(),
						},
					}
					handler.EndpointService.UpdateEndpointBySlug(result.Slug, updateData)
				}
			}
			handler.globalLock.RUnlock()
		}
	}
}
