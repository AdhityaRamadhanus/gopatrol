package grpc

import (
	"log"

	"golang.org/x/net/context"

	checkupservice "github.com/AdhityaRamadhanus/checkupd/grpc/service"
)

//ListEndpoint is grpc service that returns all the endpoints in checkup server
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
