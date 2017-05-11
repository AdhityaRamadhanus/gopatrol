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
		},
	},
}
