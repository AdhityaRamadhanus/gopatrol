package daemon

import (
	"time"

	checkup "github.com/AdhityaRamadhanus/gopatrol"
	log "github.com/Sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

//ServiceHandler is a grpc server and checkup server
type Daemon struct {
	*checkup.Checkup
	CheckInterval   time.Duration
	EndpointService checkup.EndpointService
	LoggingService  checkup.LoggingService
}

//Run the main loop for checkup to check endpoints and should be run as a goroutine
func (d *Daemon) Run() {
	for {
		timer := time.After(d.CheckInterval)
		select {
		case <-timer:
			if err := d.setCheckers(); err != nil {
				log.WithError(err).Error(ErrFailedToDoCheck)
			}
			results, err := d.Check()
			if err != nil {
				log.WithError(err).Error(ErrFailedToDoCheck)
			} else {
				if err := d.logResults(results); err != nil {
					log.WithError(err).Error(ErrFailedToStoreResult)
				}
				if err := d.checkEventsAndSync(results); err != nil {
					log.WithError(err).Error(ErrFailedToStoreResult)
				}
			}

		}
	}
}

func (d *Daemon) setCheckers() error {
	endpoints, err := d.EndpointService.GetAllEndpoints(map[string]interface{}{
		"query": bson.M{},
	})
	if err != nil {
		log.WithError(err).Error("Initialize daemon failed")
		return err
	}
	var checkers []checkup.Checker
	// Load the checkers to memory
	for _, endpoint := range endpoints {
		typeEndpoint := endpoint.(bson.M)["type"]
		switch typeEndpoint {
		case "tcp":
			var tcpChecker checkup.TCPChecker
			bsonBytes, err := bson.Marshal(endpoint)
			if err != nil {
				log.WithError(err).Error(ErrFailedToMarshalBSON)
			}
			if err := bson.Unmarshal(bsonBytes, &tcpChecker); err != nil {
				log.WithError(err).Error(ErrFailedToUnmarshalBSON)
			}
			checkers = append(checkers, tcpChecker)
		case "http":
			var httpChecker checkup.HTTPChecker
			bsonBytes, err := bson.Marshal(endpoint)
			if err != nil {
				log.WithError(err).Error(ErrFailedToMarshalBSON)
			}
			if err := bson.Unmarshal(bsonBytes, &httpChecker); err != nil {
				log.WithError(err).Error(ErrFailedToUnmarshalBSON)
			}
			checkers = append(checkers, httpChecker)
		case "dns":
			var dnsChecker checkup.DNSChecker
			bsonBytes, err := bson.Marshal(endpoint)
			if err != nil {
				log.WithError(err).Error(ErrFailedToMarshalBSON)
			}
			if err := bson.Unmarshal(bsonBytes, &dnsChecker); err != nil {
				log.WithError(err).Error(ErrFailedToUnmarshalBSON)
			}
			checkers = append(checkers, dnsChecker)
		}
	}
	d.Checkers = checkers
	return nil
}

func (d *Daemon) logResults(results []checkup.Result) error {
	for _, result := range results {
		if err := d.LoggingService.InsertLog(result); err != nil {
			log.Println("Error Inserting Log", err)
		}
	}
	return nil
}

func (d *Daemon) checkEventsAndSync(results []checkup.Result) error {
	// var events []checkup.Event
	var err error
	for _, result := range results {

		if result.Event != nil {
			log.WithFields(log.Fields{
				"message": result.Event.Message,
				"time":    result.Event.Timestamp,
				"url":     result.Event.URL,
			}).Info("New Events")
		}

		// Update Status in MONGODB to sync with the memoery

		var updateData = bson.M{
			"$set": bson.M{
				"lastchecked": result.Timestamp,
				"laststatus":  result.Status(),
			},
		}
		err = d.EndpointService.UpdateEndpointBySlug(result.Slug, updateData)
	}

	return err
}
