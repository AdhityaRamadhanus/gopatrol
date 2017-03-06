package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"

	"github.com/AdhityaRamadhanus/checkup"
	"github.com/AdhityaRamadhanus/checkup-server/checkupservice"
	"github.com/pkg/errors"

	context "golang.org/x/net/context"
)

type ServiceHandler struct {
	CheckupServer *checkup.Checkup
	globalLock    sync.RWMutex
}

func NewServiceHandler(configFile string) (*ServiceHandler, error) {
	serviceHandler := &ServiceHandler{
		CheckupServer: &checkup.Checkup{},
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

// Run as a goroutine

func (handler *ServiceHandler) runCheck() {
	for {
		select {
		case <-time.After(time.Second * 10):
			// Obtain Lock, makesure no function updating the Checkers
			handler.globalLock.RLock()
			if err := handler.CheckupServer.CheckAndStore(); err != nil {
				log.Println(err)
			}
			handler.globalLock.RUnlock()
		}
	}
}

func (handler *ServiceHandler) AddHttpEndpoint(ctx context.Context, request *checkupservice.HttpEndpointRequest) (*checkupservice.EndpointResponse, error) {
	handler.globalLock.Lock()
	defer handler.globalLock.Unlock()
	httpChecker := checkup.HTTPChecker{
		Name:     request.Endpoint.Name,
		URL:      request.Endpoint.Url,
		Attempts: 5,
	}
	handler.CheckupServer.Checkers = append(handler.CheckupServer.Checkers, httpChecker)
	message := fmt.Sprintf("Http Endpoint %s->%s Added", httpChecker.Name, httpChecker.URL)
	log.Printf(message)
	return &checkupservice.EndpointResponse{Message: message}, nil
}

func (handler *ServiceHandler) AddTcpEndpoint(ctx context.Context, request *checkupservice.TcpEndpointRequest) (*checkupservice.EndpointResponse, error) {
	handler.globalLock.Lock()
	defer handler.globalLock.Unlock()
	tcpChecker := checkup.TCPChecker{
		Name:     request.Endpoint.Name,
		URL:      request.Endpoint.Url,
		Attempts: 5,
	}
	handler.CheckupServer.Checkers = append(handler.CheckupServer.Checkers, tcpChecker)
	message := fmt.Sprintf("Tcp Endpoint %s->%s Added", tcpChecker.Name, tcpChecker.URL)
	log.Printf(message)
	return &checkupservice.EndpointResponse{Message: message}, nil
}

func (handler *ServiceHandler) DeleteEndpoint(ctx context.Context, request *checkupservice.InquiryEndpointRequest) (*checkupservice.EndpointResponse, error) {
	handler.globalLock.Lock()
	defer handler.globalLock.Unlock()
	var deletedEndpoint checkup.Checker
	for i, endpoint := range handler.CheckupServer.Checkers {
		if endpoint.GetName() == request.Name {
			deletedEndpoint = handler.CheckupServer.Checkers[i]
			handler.CheckupServer.Checkers = append(handler.CheckupServer.Checkers[:i], handler.CheckupServer.Checkers[i+1:]...)
			break
		}
	}
	message := fmt.Sprintf("Endpoint %s->%s Deleted", deletedEndpoint.GetName(), deletedEndpoint.GetURL())
	log.Printf(message)
	return &checkupservice.EndpointResponse{Message: message}, nil
}

func (handler *ServiceHandler) ListEndpoint(ctx context.Context, request *checkupservice.ListEndpointRequest) (*checkupservice.ListEndpointResponse, error) {
	handler.globalLock.RLock()
	var response = &checkupservice.ListEndpointResponse{
		Endpoints: []*checkupservice.ListEndpointResponse_Endpoint{},
	}
	var err error
	if request.Check == true {
		results, err := handler.CheckupServer.Check()
		if err != nil {
			log.Println("Error Checking", err)
		}
		for _, result := range results {
			statusMessage := "Healthy"
			if result.Down == true {
				statusMessage = "Down"
			}
			response.Endpoints = append(response.Endpoints, &checkupservice.ListEndpointResponse_Endpoint{
				Name:   result.Title,
				Url:    result.Endpoint,
				Status: statusMessage,
			})
		}
	} else {
		for _, checker := range handler.CheckupServer.Checkers {
			response.Endpoints = append(response.Endpoints, &checkupservice.ListEndpointResponse_Endpoint{
				Name: checker.GetName(),
				Url:  checker.GetURL(),
			})
		}
	}
	handler.globalLock.RUnlock()
	return response, err
}
