package grpc

import (
	"fmt"
	"log"

	"golang.org/x/net/context"

	checkup "github.com/AdhityaRamadhanus/checkupd"
	checkupservice "github.com/AdhityaRamadhanus/checkupd/grpc/service"
)

//DeleteEndpoint is grpc service to delete some endpoint
func (handler *ServiceHandler) DeleteEndpoint(ctx context.Context, request *checkupservice.DeleteEndpointRequest) (*checkupservice.EndpointResponse, error) {
	handler.globalLock.Lock()
	defer handler.globalLock.Unlock()
	var deletedEndpoint checkup.Checker
	for i, endpoint := range handler.CheckupServer.Checkers {
		if endpoint.GetURL() == request.Url {
			deletedEndpoint = handler.CheckupServer.Checkers[i]
			handler.CheckupServer.Checkers = append(handler.CheckupServer.Checkers[:i], handler.CheckupServer.Checkers[i+1:]...)
			break
		}
	}
	message := "Cannot find endpoint to delete"
	if deletedEndpoint != nil {
		message = fmt.Sprintf("Endpoint %s->%s Deleted", deletedEndpoint.GetName(), deletedEndpoint.GetURL())
	}
	log.Printf(message)
	return &checkupservice.EndpointResponse{Message: message}, nil
}
