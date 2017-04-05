package commands

import (
	"context"
	"log"

	checkupservice "github.com/AdhityaRamadhanus/checkupd/grpc/service"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

func addEndpoint(cliContext *cli.Context) {
	conn, err := grpc.Dial(cliContext.String("host"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := checkupservice.NewCheckupClient(conn)

	r, err := c.AddHttpEndpoint(context.Background(), &checkupservice.HttpEndpointRequest{
		Endpoint: &checkupservice.GenericEndpointRequest{
			Name:     cliContext.String("name"),
			Url:      cliContext.String("url"),
			Attempts: 5,
		},
	})

	if err != nil {
		log.Fatalf("Could not add endpoint", err)
	}
	log.Println(r.Message)
}
