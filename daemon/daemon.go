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
	CheckersService checkup.CheckersService
	EventService    checkup.EventService
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
				// if err := d.logResults(results); err != nil {
				// 	log.WithError(err).Error(ErrFailedToStoreResult)
				// }
				if err := d.checkResultsAndSync(results); err != nil {
					log.WithError(err).Error(ErrFailedToStoreResult)
				}
			}

		}
	}
}

func (d *Daemon) setCheckers() error {
	endpoints, err := d.CheckersService.GetAllCheckers(map[string]interface{}{
		"query":      bson.M{},
		"pagination": false,
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

func (d *Daemon) checkResultsAndSync(results []checkup.Result) error {
	var err error
	for _, result := range results {
		updateData := bson.M{
			"lastchecked": result.Timestamp,
			"laststatus":  result.Status(),
		}

		if result.Event {
			event := checkup.NewEvent(result)

			log.WithFields(log.Fields{
				"message": event.Message,
				"time":    event.Timestamp,
				"url":     event.URL,
			}).Info("New Events")
			// Ignore the error for now
			d.EventService.InsertEvent(event)
		}

		if result.Notification {
			updateData["lastchange"] = result.Timestamp
			log.WithFields(log.Fields{
				"message": result.URL + " is " + result.Status(),
				"time":    result.Timestamp,
				"url":     result.URL,
			}).Info("New Notification")
			for _, notifier := range d.Notifier {
				// ignore the error, for now
				go notifier.Notify(result)
			}
		}

		// Update Status in MONGODB to sync with the memoery

		err = d.CheckersService.UpdateCheckerBySlug(result.Slug, bson.M{
			"$set": updateData,
		})
	}

	return err
}
