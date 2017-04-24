package daemon

import (
	cli "github.com/urfave/cli"
)

// Commands is a list of commands that will be used in main function of cli app
var Commands = cli.Commands{
	{
		Name:   "daemon",
		Usage:  "run daemon",
		Action: runDaemon,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "log",
				Value: "gopatrol-daemon.log",
				Usage: "",
			},
			cli.StringFlag{
				Name:  "config",
				Value: "checkup_config/checkup.json",
				Usage: "config file for checkup",
			},
			cli.StringFlag{
				Name:  "proto",
				Value: "unix",
				Usage: "protocol to run the daemon",
			},
			cli.StringFlag{
				Name:  "address",
				Value: "/tmp/gopatrol.sock",
				Usage: "address of the daemon",
			},
		},
	},
	{
		Name:   "tls-daemon",
		Usage:  "run tls daemon",
		Action: runTLSDaemon,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "log",
				Value: "gopatrol-status.log",
				Usage: "",
			},
			cli.StringFlag{
				Name:  "config",
				Value: "checkup_config/checkup.json",
				Usage: "config file for checkup",
			},
			cli.StringFlag{
				Name:  "address",
				Value: ":9009",
				Usage: "address of the daemon",
			},
			cli.StringFlag{
				Name:  "cert",
				Value: "",
				Usage: "certificate file",
			},
			cli.StringFlag{
				Name:  "key",
				Value: "",
				Usage: "key file",
			},
		},
	},
	{
		Name:  "setup",
		Usage: "Setup Daemon and create status page",
		Subcommands: []cli.Command{
			{
				Name:   "fs",
				Usage:  "setup daemon and status page using filesystem as storage",
				Action: SetupFS,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "url",
						Value: "localhost:9008",
						Usage: "Url to serve the status page, default to localhost:9008",
					},
				},
			},
			{
				Name:   "s3",
				Usage:  "setup daemon and status page using s3 as storage",
				Action: SetupS3,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "url",
						Value: "localhost:9008",
						Usage: "Url to serve the status page, default to localhost:9008",
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
	{
		Name:   "status-page",
		Usage:  "Serve status page",
		Action: serveStatusPage,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "type",
				Value: "http",
				Usage: "Url to serve the status page, default to localhost:80",
			},
			cli.StringFlag{
				Name:  "config",
				Value: "caddy_config/Caddyfile",
				Usage: "Caddyfile to run",
			},
			cli.StringFlag{
				Name:  "log",
				Value: "gopatrol-status.log",
				Usage: "",
			},
		},
	},
}
