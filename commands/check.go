package commands

import (
	"context"
	"log"

	checkupservice "github.com/AdhityaRamadhanus/checkupd/grpc/service"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

func checkEndpoint(cliContext *cli.Context) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(cliContext.String("host"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := checkupservice.NewCheckupClient(conn)

	r, err := c.ListEndpoint(context.Background(), &checkupservice.ListEndpointRequest{Check: true})

	if err != nil {
		log.Fatalf("could not list endpoints: %v", err)
	} else {
		for _, endpoint := range r.Endpoints {
			log.Println(endpoint.Name, " ", endpoint.Url, " ", endpoint.Status)
		}
	}
}
