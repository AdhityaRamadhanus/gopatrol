package grpc

import (
	"log"
	"sync"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	checkup "github.com/AdhityaRamadhanus/gopatrol"
	"github.com/AdhityaRamadhanus/gopatrol/mongo"
)

//ServiceHandler is a grpc server and checkup server
type ServiceHandler struct {
	CheckupServer   *checkup.Checkup
	globalLock      sync.RWMutex
	CheckInterval   time.Duration
	EndpointService *mongo.EndpointService
}

//NewServiceHandler create new ServiceHandler from a configfile (checkup.json)
func NewServiceHandler(mongoSession *mgo.Session) (*ServiceHandler, error) {
	serviceHandler := &ServiceHandler{
		CheckupServer: &checkup.Checkup{},
	}
	serviceHandler.EndpointService = mongo.NewEndpointService(mongoSession)
	endpoints, _ := serviceHandler.EndpointService.GetAllEndpoints(bson.M{}, 0, 10)
	for _, endpoint := range endpoints {
		typeEndpoint := endpoint.(bson.M)["type"]
		switch typeEndpoint {
		case "tcp":
			var tcpChecker checkup.TCPChecker
			bsonBytes, _ := bson.Marshal(endpoint)
			bson.Unmarshal(bsonBytes, &tcpChecker)
			serviceHandler.CheckupServer.Checkers = append(serviceHandler.CheckupServer.Checkers, tcpChecker)
		case "http":
			var httpChecker checkup.HTTPChecker
			bsonBytes, _ := bson.Marshal(endpoint)
			bson.Unmarshal(bsonBytes, &httpChecker)
			serviceHandler.CheckupServer.Checkers = append(serviceHandler.CheckupServer.Checkers, httpChecker)
		case "dns":
			var dnsChecker checkup.DNSChecker
			bsonBytes, _ := bson.Marshal(endpoint)
			bson.Unmarshal(bsonBytes, &dnsChecker)
			serviceHandler.CheckupServer.Checkers = append(serviceHandler.CheckupServer.Checkers, dnsChecker)
		}
	}
	serviceHandler.CheckupServer.Storage = mongo.NewLoggingService(mongoSession)

	return serviceHandler, nil
}

//Run the main loop for checkup to check endpoints and should be run as a goroutine
func (handler *ServiceHandler) Run() {
	for {
		timer := time.After(handler.CheckInterval)
		select {
		case <-timer: //hardcoded for now
			// Obtain Lock, makesure no function updating the Checkers
			handler.globalLock.RLock()
			results, err := handler.CheckupServer.Check()
			if err != nil {
				log.Println(err)
			} else {
				if err := handler.CheckupServer.Storage.Store(results); err != nil {
					log.Println(err)
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
