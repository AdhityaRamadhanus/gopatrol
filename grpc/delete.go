package grpc

import (
	"fmt"
	"log"

	"golang.org/x/net/context"

	"github.com/AdhityaRamadhanus/checkup"
	checkupservice "github.com/AdhityaRamadhanus/checkupd/grpc/service"
)

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
