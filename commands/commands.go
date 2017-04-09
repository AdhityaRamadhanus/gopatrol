package commands

import (
	cli "github.com/urfave/cli"
)

// Commands is a list of commands that will be used in main function of cli app
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
		Name:  "setup-page",
		Usage: "Setup statuspage",
		Subcommands: []cli.Command{
			{
				Name:   "fs",
				Usage:  "setup status page using filesystem as storage",
				Action: setupFSStatusPage,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "url",
						Value: "localhost:80",
						Usage: "url to serve the status page",
					},
				},
			},
			{
				Name:   "s3",
				Usage:  "setup status page using filesystem as storage",
				Action: setupS3StatusPage,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "url",
						Value: "localhost:80",
						Usage: "url to serve the status page",
					},
					cli.StringFlag{
						Name:  "accesskeyid, i",
						Value: "",
						Usage: "S3 Access Key ID",
					},
					cli.StringFlag{
						Name:  "secretaccesskey, k",
						Value: "",
						Usage: "S3 Secret",
					},
					cli.StringFlag{
						Name:  "region, r",
						Value: "",
						Usage: "S3 Region",
					},
					cli.StringFlag{
						Name:  "bucket, b",
						Value: "",
						Usage: "S3 Bucket",
					},
				},
			},
		},
	},
}
