package commands

import (
	cli "github.com/urfave/cli"
)

var Commands = cli.Commands{
	{
		Name:   "add-http",
		Usage:  "Add endpoints to checkup",
		Action: addHTTPEndpoint,
		Flags: []cli.Flag{
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
	{
		Name:   "setup-page",
		Usage:  "delete endpoint",
		Action: setupStatusPage,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "url",
				Value: "localhost:80",
				Usage: "url to serve the status page",
			},
			cli.StringFlag{
				Name:  "storage",
				Value: "fs",
				Usage: "storage type",
			},
		},
	},
}
