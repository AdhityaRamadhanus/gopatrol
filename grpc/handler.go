package grpc

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"

	"github.com/AdhityaRamadhanus/checkup"
	"github.com/pkg/errors"
)

//ServiceHandler is a grpc server and checkup server
type ServiceHandler struct {
	CheckupServer *checkup.Checkup
	globalLock    sync.RWMutex
	ConfigPath    string
}

//NewServiceHandler create new ServiceHandler from a configfile (checkup.json)
func NewServiceHandler(configFile string) (*ServiceHandler, error) {
	serviceHandler := &ServiceHandler{
		CheckupServer: &checkup.Checkup{},
		ConfigPath:    configFile,
	}
	file, err := os.OpenFile(configFile, os.O_RDONLY, 777)
	if err != nil {
		return nil, errors.Wrap(err, "Error loading config file")
	}
	configBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.Wrap(err, "Error reading config file")
	}
	if err := serviceHandler.CheckupServer.UnmarshalJSON(configBytes); err != nil {
		return nil, errors.Wrap(err, "Error unmarshal config file")
	}

	return serviceHandler, nil
}

//Run the main loop for checkup to check endpoints and should be run as a goroutine
func (handler *ServiceHandler) Run() {
	for {
		timer := time.After(time.Second * 10)
		select {
		case <-timer: //hardcoded for now
			// Obtain Lock, makesure no function updating the Checkers
			handler.globalLock.RLock()
			if err := handler.CheckupServer.CheckAndStore(); err != nil {
				log.Println(err)
			}
			handler.globalLock.RUnlock()
		}
	}
}

// SerializeJSON is a function that save current checkup config to checkup.json in current working dir
func (handler *ServiceHandler) SerializeJSON() error {
	jsonBytes, err := handler.CheckupServer.MarshalJSON()
	if err != nil {
		return err
	}
	file, err := os.OpenFile(handler.ConfigPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	defer file.Close()
	if err != nil {
		return errors.Wrap(err, "Failed opening file")
	}
	_, err = file.Write(jsonBytes)
	return errors.Wrap(err, "Failed to write json")
}
