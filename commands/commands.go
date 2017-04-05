package commands

import (
	cli "github.com/urfave/cli"
)

var Commands = cli.Commands{
	{
		Name:   "add",
		Usage:  "Add endpoints to checkup",
		Action: addEndpoint,
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
}
