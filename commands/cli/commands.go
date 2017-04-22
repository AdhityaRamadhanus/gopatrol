package cli

import (
	"time"

	cli "github.com/urfave/cli"
)

var defaultConnFlags = []cli.Flag{
	cli.BoolFlag{
		Name:  "tls",
		Usage: "Send request over tls",
	},
	cli.StringFlag{
		Name:  "host",
		Value: "/tmp/gopatrol.sock",
		Usage: "grpc server address",
	},
}

var defaultEndpointFlags = []cli.Flag{
	cli.IntFlag{
		Name:  "attempts, a",
		Value: 5,
		Usage: "how many times to check endpoint",
	},
	cli.Int64Flag{
		Name:  "thresholdrtt, rtt",
		Value: 0,
		Usage: "Threshold Rtt to define a degraded endpoint",
	},
}

// Commands is a list of commands that will be used in main function of cli app
var Commands = cli.Commands{
	{
		Name:      "add-http",
		Usage:     "Add endpoints to checkup",
		Action:    addHTTPEndpoint,
		UsageText: "checklist add-http [command options] name url",
		Flags: append(
			append(defaultConnFlags, defaultEndpointFlags...),
			cli.StringFlag{
				Name:  "mustcontain",
				Value: "",
				Usage: "HTML content that a page should contain to determine whether a page is up or down",
			},
			cli.StringFlag{
				Name:  "mustnotcontain",
				Value: "",
				Usage: "HTML content that a page should not contain to determine whether a page is up or down",
			},
			cli.StringFlag{
				Name:  "headers",
				Value: "",
				Usage: "Http Headers to send along the check request",
			},
			cli.IntFlag{
				Name:  "upstatus",
				Value: 200,
				Usage: "Http status code to define a healthy page",
			},
		),
	},
	{
		Name:      "add-tcp",
		Usage:     "Add tcp endpoints to checkup",
		Action:    addTCPEndpoint,
		UsageText: "checklist add-http [command options] name url",
		Flags: append(
			append(defaultConnFlags, defaultEndpointFlags...),
			cli.BoolFlag{
				Name:  "tls-enabled",
				Usage: "Enable TLS connection to endpoint",
			},
			cli.StringFlag{
				Name:  "tls-ca",
				Value: "",
				Usage: "Certificate file to established tls connection",
			},
			cli.BoolFlag{
				Name:  "tls-skip-verify",
				Usage: "Skip verify tls certificate",
			},
			cli.Int64Flag{
				Name:  "timeout",
				Value: int64(3 * time.Second),
				Usage: "Timeout to established a tls connection",
			},
		),
	},
	{
		Name:      "add-dns",
		Usage:     "Add dns endpoints to checkup",
		Action:    addDNSEndpoint,
		UsageText: "checklist add-dns [command options] name url host",
		Flags: append(
			append(defaultConnFlags, defaultEndpointFlags...),
			cli.Int64Flag{
				Name:  "timeout",
				Value: int64(3 * time.Second),
				Usage: "Timeout to established a tls connection",
			},
		),
	},
	{
		Name:   "check",
		Usage:  "list and check endpoints",
		Action: checkEndpoint,
		Flags:  defaultConnFlags,
	},
	{
		Name:   "list",
		Usage:  "list endpoint",
		Action: listEndpoint,
		Flags:  defaultConnFlags,
	},
	{
		Name:   "delete",
		Usage:  "delete endpoint",
		Action: deleteEndpoint,
		Flags:  defaultConnFlags,
	},
}
