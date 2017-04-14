package checklist

import (
	"context"
	"log"

	checkupservice "github.com/AdhityaRamadhanus/checkupd/grpc/service"
	"github.com/urfave/cli"
)

func deleteEndpoint(cliContext *cli.Context) {
	// Set up a connection to the server.
	conn, err := createGrpcClient(cliContext)
	if err != nil {
		log.Println("could not connect to grpc server", err)
		return
	}
	defer conn.Close()
	c := checkupservice.NewCheckupClient(conn)

	r, err := c.DeleteEndpoint(context.Background(), &checkupservice.InquiryEndpointRequest{Name: cliContext.String("name")})

	if err != nil {
		log.Println("failed to delete endpoint", err)
	} else {
		log.Println(r.Message)
	}
}
