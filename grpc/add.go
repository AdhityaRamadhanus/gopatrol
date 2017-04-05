package grpc

import (
	"fmt"
	"log"

	"golang.org/x/net/context"

	"github.com/AdhityaRamadhanus/checkup"
	checkupservice "github.com/AdhityaRamadhanus/checkupd/grpc/service"
)

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
