package checklist

import (
	"context"
	"log"

	checkupservice "github.com/AdhityaRamadhanus/checkupd/grpc/service"
	"github.com/urfave/cli"
)

func listEndpoint(cliContext *cli.Context) {
	// Set up a connection to the server.
	conn, err := createGrpcClient(cliContext)
	if err != nil {
		log.Println("could not connect to grpc server", err)
		return
	}
	defer conn.Close()
	c := checkupservice.NewCheckupClient(conn)

	r, err := c.ListEndpoint(context.Background(), &checkupservice.ListEndpointRequest{Check: false})

	if err != nil {
		log.Println("failed to get list of endpoints", err)
	} else {
		for _, endpoint := range r.Endpoints {
			log.Println(endpoint.Name, " ", endpoint.Url, " ", endpoint.Status)
		}
	}
}
