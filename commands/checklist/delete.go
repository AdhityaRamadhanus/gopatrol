package checklist

import (
	"context"
	"fmt"
	"log"

	checkupservice "github.com/AdhityaRamadhanus/checkupd/grpc/service"
	"github.com/urfave/cli"
)

func deleteEndpoint(c *cli.Context) error {
	endpointURL := ""
	if c.NArg() > 0 {
		endpointURL = c.Args().Get(0)
	} else {
		fmt.Println("Please provice endpoint url to delete")
		cli.ShowCommandHelp(c, "delete")
		return cli.NewExitError("", 1)
	}
	conn, err := createGrpcClient(c)
	if err != nil {
		errMessage := "Couldn't connect to grpc server: " + err.Error()
		return cli.NewExitError(errMessage, 1)
	}
	defer conn.Close()
	service := checkupservice.NewCheckupClient(conn)

	r, err := service.DeleteEndpoint(context.Background(), &checkupservice.DeleteEndpointRequest{
		Url: endpointURL,
	})
	if err != nil {
		errMessage := "Failed to delete endpoint :" + err.Error()
		return cli.NewExitError(errMessage, 1)
	}

	log.Println(r.Message)
	return nil
}
