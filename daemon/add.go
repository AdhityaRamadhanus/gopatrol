package daemon

import (
	checkup "github.com/AdhityaRamadhanus/gopatrol"
	log "github.com/Sirupsen/logrus"
)

//AddTCPEndpoint is grpc service to add tcp endpoint to checkup server
func (handler *Daemon) AddEndpoint(checker checkup.Checker) error {
	if err := handler.EndpointService.InsertEndpoint(checker); err != nil {
		log.WithError(err).Error("Error inserting endpoint")
		return err
	}
	handler.globalLock.Lock()
	defer handler.globalLock.Unlock()
	handler.Checkers = append(handler.Checkers, checker)
	log.WithFields(log.Fields{
		"name": checker.GetName(),
		"url":  checker.GetURL(),
	}).Info("Add Endpoint")
	return nil
}
