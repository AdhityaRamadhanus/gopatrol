package checklist

import (
	"os"

	cli "github.com/urfave/cli"
)

// Commands is a list of commands that will be used in main function of cli app
var Commands = cli.Commands{
	{
		Name:   "add-http",
		Usage:  "Add endpoints to checkup",
		Action: addHTTPEndpoint,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "tls",
				Usage: "Name of endpoint",
			},
			cli.StringFlag{
				Name:  "name",
				Value: "",
				Usage: "Name of endpoint",
			},
			cli.StringFlag{
				Name:  "url",
				Value: "",
				Usage: "URL to check",
			},
			cli.StringFlag{
				Name:  "host",
				Value: ":9009",
				Usage: "grpc server address",
			},
		},
	},
	{
		Name:   "add-tcp",
		Usage:  "Add tcp endpoints to checkup",
		Action: addTCPEndpoint,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "tls",
				Usage: "Name of endpoint",
			},
			cli.StringFlag{
				Name:  "name",
				Value: "",
				Usage: "Name of endpoint",
			},
			cli.StringFlag{
				Name:  "address",
				Value: "",
				Usage: "Address to check",
			},
			cli.BoolFlag{
				Name:  "tls",
				Usage: "Is it tls endpoint?",
			},
			cli.StringFlag{
				Name:  "host",
				Value: ":9009",
				Usage: "grpc server address",
			},
		},
	},
	{
		Name:   "check",
		Usage:  "list and check endpoints",
		Action: checkEndpoint,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "tls",
				Usage: "Name of endpoint",
			},
			cli.StringFlag{
				Name:  "pass, p",
				Value: os.Getenv("CHECKUPD_CLIENT_PASS"),
				Usage: "Password for grpc call",
			},
			cli.StringFlag{
				Name:  "host",
				Value: ":9009",
				Usage: "grpc server address",
			},
		},
	},
	{
		Name:   "list",
		Usage:  "list endpoint",
		Action: listEndpoint,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "tls",
				Usage: "Name of endpoint",
			},
			cli.StringFlag{
				Name:  "host",
				Value: ":9009",
				Usage: "grpc server address",
			},
		},
	},
	{
		Name:   "delete",
		Usage:  "delete endpoint",
		Action: deleteEndpoint,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "tls",
				Usage: "Name of endpoint",
			},
			cli.StringFlag{
				Name:  "pass, p",
				Value: os.Getenv("CHECKUPD_CLIENT_PASS"),
				Usage: "Password for grpc call",
			},
			cli.StringFlag{
				Name:  "host",
				Value: ":9009",
				Usage: "grpc server address",
			},
			cli.StringFlag{
				Name:  "name",
				Value: "",
				Usage: "endpoint name",
			},
		},
	},
}
