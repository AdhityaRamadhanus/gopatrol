package daemon

import (
	"os"

	log "github.com/Sirupsen/logrus"
	cli "github.com/urfave/cli"
)

func fatalError(err error) {
	log.WithError(err).Error("FATAL ERROR")
	os.Exit(1)
}

// Commands is a list of commands that will be used in main function of cli app
var Commands = cli.Commands{
	{
		Name:   "daemon",
		Usage:  "run gopatrol checking daemon",
		Action: runDaemon,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "log",
				Value: "gopatrol-daemon.log",
				Usage: "",
			},
		},
	},
	{
		Name:   "api",
		Usage:  "run gopatrol api server",
		Action: runApiServer,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "log",
				Value: "gopatrol-daemon.log",
				Usage: "",
			},
		},
	},
}
