package checklist

import (
	"context"
	"log"

	checkupservice "github.com/AdhityaRamadhanus/checkupd/grpc/service"
	"github.com/urfave/cli"
)

func deleteEndpoint(c *cli.Context) error {
	conn, err := createGrpcClient(c)
	if err != nil {
		errMessage := "Couldn't connect to grpc server: " + err.Error()
		return cli.NewExitError(errMessage, 1)
	}
	defer conn.Close()
	service := checkupservice.NewCheckupClient(conn)

	r, err := service.DeleteEndpoint(context.Background(), &checkupservice.DeleteEndpointRequest{
		Url: c.String("url"),
	})
	if err != nil {
		errMessage := "Failed to delete endpoint :" + err.Error()
		return cli.NewExitError(errMessage, 1)
	}

	log.Println(r.Message)
	return nil
}
