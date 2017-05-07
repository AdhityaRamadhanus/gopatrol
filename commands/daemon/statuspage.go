package daemon

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/mholt/caddy"
	_ "github.com/mholt/caddy/caddyhttp"
	"github.com/urfave/cli"
	"github.com/xenolf/lego/acme"
	"gopkg.in/natefinch/lumberjack.v2"
)

const appName = "gopatrol"
const appVersion = "1.0.0"

// Run is Caddy's main() function.
func serveStatusPage(c *cli.Context) error {
	caddy.TrapSignals()
	caddy.RegisterCaddyfileLoader("flag", caddy.LoaderFunc(func(serverType string) (caddy.Input, error) {
		contents, err := ioutil.ReadFile(c.String("config"))
		if err != nil {
			return nil, err
		}
		return caddy.CaddyfileInput{
			Contents:       contents,
			Filepath:       c.String("config"),
			ServerTypeName: serverType,
		}, nil
	}))

	caddy.AppName = appName
	caddy.AppVersion = appVersion
	acme.UserAgent = appName + "/" + appVersion

	// Set up process log before anything bad happens
	switch c.String("log") {
	case "stdout":
		log.SetOutput(os.Stdout)
	case "stderr":
		log.SetOutput(os.Stderr)
	default:
		log.SetOutput(&lumberjack.Logger{
			Filename:   c.String("log"),
			MaxSize:    100,
			MaxAge:     14,
			MaxBackups: 10,
		})
	}

	// Executes Startup events
	// caddy.
	// caddy.EmitEvent(caddy.StartupEvent)

	// Get Caddyfile input
	caddyfileinput, err := caddy.LoadCaddyfile(c.String("type"))
	if err != nil {
		return err
	}

	// Start your engines
	instance, err := caddy.Start(caddyfileinput)
	if err != nil {
		return err
	}

	// Twiddle your thumbs
	instance.Wait()
	return nil
}
