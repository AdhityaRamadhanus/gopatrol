package checklist

import (
	"context"
	"log"

	checkupservice "github.com/AdhityaRamadhanus/checkupd/grpc/service"
	"github.com/urfave/cli"
)

func addHTTPEndpoint(cliContext *cli.Context) {
	conn, err := createGrpcClient(cliContext)
	if err != nil {
		log.Println("could not connect to grpc server", err)
		return
	}
	defer conn.Close()
	c := checkupservice.NewCheckupClient(conn)

	r, err := c.AddHTTPEndpoint(context.Background(), &checkupservice.HttpEndpointRequest{
		Endpoint: &checkupservice.GenericEndpointRequest{
			Name:     cliContext.String("name"),
			Url:      cliContext.String("url"),
			Attempts: 5,
		},
	})

	if err != nil {
		log.Println("Could not add endpoint", err)
	} else {
		log.Println(r.Message)
	}
}

func addTCPEndpoint(cliContext *cli.Context) {
	conn, err := createGrpcClient(cliContext)
	if err != nil {
		log.Println("could not connect to grpc server", err)
		return
	}
	defer conn.Close()
	c := checkupservice.NewCheckupClient(conn)

	r, err := c.AddTCPEndpoint(context.Background(), &checkupservice.TcpEndpointRequest{
		Endpoint: &checkupservice.GenericEndpointRequest{
			Name:     cliContext.String("name"),
			Url:      cliContext.String("address"),
			Attempts: 5,
		},
		Tls: cliContext.Bool("tcp-tls"),
	})

	if err != nil {
		log.Println("Could not add endpoint", err)
	} else {
		log.Println(r.Message)
	}
}
