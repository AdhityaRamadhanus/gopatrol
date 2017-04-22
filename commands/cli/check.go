package cli

import (
	"context"
	"log"

	checkupservice "github.com/AdhityaRamadhanus/gopatrol/grpc/service"
	"github.com/urfave/cli"
)

func checkEndpoint(c *cli.Context) error {
	conn, err := createGrpcClient(c)
	if err != nil {
		errMessage := "Couldn't connect to grpc server: " + err.Error()
		return cli.NewExitError(errMessage, 1)
	}
	defer conn.Close()
	service := checkupservice.NewCheckupClient(conn)

	r, err := service.ListEndpoint(context.Background(), &checkupservice.ListEndpointRequest{Check: true})

	if err != nil {
		errMessage := "Failed to get endpoint :" + err.Error()
		return cli.NewExitError(errMessage, 1)
	}

	for _, endpoint := range r.Endpoints {
		log.Println(endpoint.Name, " ", endpoint.Url, " ", endpoint.Status)
	}
	return nil
}
