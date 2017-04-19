package checklist

import (
	"context"
	"fmt"
	"log"

	checkupservice "github.com/AdhityaRamadhanus/checkupd/grpc/service"
	"github.com/urfave/cli"
)

func addHTTPEndpoint(c *cli.Context) error {
	endpointName := ""
	endpointURL := ""
	if c.NArg() > 1 {
		endpointName = c.Args().Get(0)
		endpointURL = c.Args().Get(1)
	} else {
		fmt.Println("Please provice name and url to add http endpoint")
		cli.ShowCommandHelp(c, "add-http")
		return cli.NewExitError("", 1)
	}

	request := &checkupservice.AddHttpEndpointRequest{
		Endpoint: &checkupservice.GenericEndpointRequest{
			Name:         endpointName,
			Url:          endpointURL,
			Attempts:     int32(c.Int("attempts")),
			Thresholdrtt: c.Int64("thresholdrtt"),
		},
		MustContain:    c.String("mustcontain"),
		MustNotContain: c.String("mustnotcontain"),
		Headers:        c.String("headers"),
		Upstatus:       int32(c.Int("upstatus")),
	}

	conn, err := createGrpcClient(c)
	if err != nil {
		errMessage := "Couldn't connect to grpc server: " + err.Error()
		return cli.NewExitError(errMessage, 1)
	}
	defer conn.Close()
	service := checkupservice.NewCheckupClient(conn)

	r, err := service.AddHTTPEndpoint(context.Background(), request)

	if err != nil {
		errMessage := "Failed to add http endpoint :" + err.Error()
		return cli.NewExitError(errMessage, 1)
	}
	log.Println(r.Message)
	return nil
}

func addTCPEndpoint(c *cli.Context) error {
	endpointName := ""
	endpointURL := ""
	if c.NArg() > 1 {
		endpointName = c.Args().Get(0)
		endpointURL = c.Args().Get(1)
	} else {
		fmt.Println("Please provice name and url to add tcp endpoint")
		cli.ShowCommandHelp(c, "add-tcp")
		return cli.NewExitError("", 1)
	}

	request := &checkupservice.AddTcpEndpointRequest{
		Endpoint: &checkupservice.GenericEndpointRequest{
			Name:         endpointName,
			Url:          endpointURL,
			Attempts:     int32(c.Int("attempts")),
			Thresholdrtt: c.Int64("thresholdrtt"),
		},
		TlsEnabled:    c.Bool("tls-enabled"),
		TlsCaFile:     c.String("tls-ca"),
		TlsSkipVerify: c.Bool("tls-skip-verify"),
		Timeout:       c.Int64("timeout"),
	}

	conn, err := createGrpcClient(c)
	if err != nil {
		errMessage := "Couldn't connect to grpc server: " + err.Error()
		return cli.NewExitError(errMessage, 1)
	}
	defer conn.Close()
	service := checkupservice.NewCheckupClient(conn)

	r, err := service.AddTCPEndpoint(context.Background(), request)

	if err != nil {
		errMessage := "Failed to add tcp endpoint :" + err.Error()
		return cli.NewExitError(errMessage, 1)
	}
	log.Println(r.Message)
	return nil
}
