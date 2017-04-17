package checkupd

import (
	cli "github.com/urfave/cli"
)

// Commands is a list of commands that will be used in main function of cli app
var Commands = cli.Commands{
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
	{
		Name:  "setup-daemon",
		Usage: "Setup daemon",
		Subcommands: []cli.Command{
			{
				Name:   "fs",
				Usage:  "setup status page using filesystem as storage",
				Action: setupFSDaemon,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "log",
						Value: "./checkup_config/logs",
						Usage: "url to serve the status page",
					},
				},
			},
			{
				Name:   "s3",
				Usage:  "setup status page using filesystem as storage",
				Action: setupS3Daemon,
				Flags: []cli.Flag{
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
		Name:   "daemon",
		Usage:  "run daemon",
		Action: runDaemon,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "config",
				Value: "checkup_config/checkup.json",
				Usage: "url to bind the daemon",
			},
			cli.StringFlag{
				Name:  "interval",
				Value: "1m",
				Usage: "url to bind the daemon",
			},
			cli.StringFlag{
				Name:  "proto",
				Value: "unix",
				Usage: "url to bind the daemon",
			},
			cli.StringFlag{
				Name:  "address",
				Value: "/tmp/checkupd.sock",
				Usage: "url to bind the daemon",
			},
		},
	},
	{
		Name:   "tls-daemon",
		Usage:  "run tls daemon",
		Action: runTLSDaemon,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "config",
				Value: "checkup_config/checkup.json",
				Usage: "url to bind the daemon",
			},
			cli.StringFlag{
				Name:  "interval",
				Value: "1m",
				Usage: "url to bind the daemon",
			},
			cli.StringFlag{
				Name:  "port",
				Value: ":9009",
				Usage: "url to bind the daemon",
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
}
