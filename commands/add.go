package commands

import (
	"context"
	"log"

	checkupservice "github.com/AdhityaRamadhanus/checkupd/grpc/service"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

func addHTTPEndpoint(cliContext *cli.Context) {
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

func addTCPEndpoint(cliContext *cli.Context) {
	conn, err := grpc.Dial(cliContext.String("host"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := checkupservice.NewCheckupClient(conn)

	r, err := c.AddTcpEndpoint(context.Background(), &checkupservice.TcpEndpointRequest{
		Endpoint: &checkupservice.GenericEndpointRequest{
			Name:     cliContext.String("name"),
			Url:      cliContext.String("address"),
			Attempts: 5,
		},
		Tls: cliContext.Bool("tls"),
	})

	if err != nil {
		log.Fatalf("Could not add endpoint", err)
	}
	log.Println(r.Message)
}
