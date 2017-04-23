package daemon

import (
	cli "github.com/urfave/cli"
)

// Commands is a list of commands that will be used in main function of cli app
var Commands = cli.Commands{
	// {
	// 	Name:  "setup-page",
	// 	Usage: "Setup statuspage",
	// 	Subcommands: []cli.Command{
	// 		{
	// 			Name:   "fs",
	// 			Usage:  "setup status page using filesystem as storage",
	// 			Action: setupFSStatusPage,
	// 			Flags: []cli.Flag{
	// 				cli.StringFlag{
	// 					Name:  "url",
	// 					Value: "localhost:80",
	// 					Usage: "url to serve the status page",
	// 				},
	// 			},
	// 		},
	// 		{
	// 			Name:   "s3",
	// 			Usage:  "setup status page using filesystem as storage",
	// 			Action: setupS3StatusPage,
	// 			Flags: []cli.Flag{
	// 				cli.StringFlag{
	// 					Name:  "url",
	// 					Value: "localhost:80",
	// 					Usage: "url to serve the status page",
	// 				},
	// 				cli.StringFlag{
	// 					Name:  "accesskeyid, i",
	// 					Value: "",
	// 					Usage: "S3 Access Key ID",
	// 				},
	// 				cli.StringFlag{
	// 					Name:  "secretaccesskey, k",
	// 					Value: "",
	// 					Usage: "S3 Secret",
	// 				},
	// 				cli.StringFlag{
	// 					Name:  "region, r",
	// 					Value: "",
	// 					Usage: "S3 Region",
	// 				},
	// 				cli.StringFlag{
	// 					Name:  "bucket, b",
	// 					Value: "us-east-1",
	// 					Usage: "S3 Bucket",
	// 				},
	// 			},
	// 		},
	// 	},
	// },
	// {
	// Name:  "setup-daemon",
	// Usage: "Setup daemon",
	// Subcommands: []cli.Command{
	// 	{
	// 		Name:   "fs",
	// 		Usage:  "setup daemon using filesystem as storage",
	// 		Action: setupFSDaemon,
	// 		Flags: []cli.Flag{
	// 			cli.StringFlag{
	// 				Name:  "log",
	// 				Value: "./checkup_config/logs",
	// 				Usage: "directory to save the logs (check results)",
	// 			},
	// 		},
	// 	},
	// 	{
	// 		Name:   "s3",
	// 		Usage:  "setup daemon using s3 as storage",
	// 		Action: setupS3Daemon,
	// 		Flags: []cli.Flag{
	// 			cli.StringFlag{
	// 				Name:  "accesskeyid, i",
	// 				Value: "",
	// 				Usage: "S3 Access Key ID",
	// 			},
	// 			cli.StringFlag{
	// 				Name:  "secretaccesskey, k",
	// 				Value: "",
	// 				Usage: "S3 Secret",
	// 			},
	// 			cli.StringFlag{
	// 				Name:  "region, r",
	// 				Value: "",
	// 				Usage: "S3 Region",
	// 			},
	// 			cli.StringFlag{
	// 				Name:  "bucket, b",
	// 				Value: "",
	// 				Usage: "S3 Bucket",
	// 			},
	// 		},
	// 	},
	// },
	// },
	{
		Name:   "daemon",
		Usage:  "run daemon",
		Action: runDaemon,
		Flags: []cli.Flag{
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
						Value: "localhost:80",
						Usage: "Url to serve the status page, default to localhost:80",
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
						Value: "localhost:80",
						Usage: "Url to serve the status page, default to localhost:80",
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
