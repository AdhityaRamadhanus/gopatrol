package daemon

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AdhityaRamadhanus/gopatrol"
	daemon "github.com/AdhityaRamadhanus/gopatrol/daemon"
	"github.com/AdhityaRamadhanus/gopatrol/mongo"
	log "github.com/Sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"

	"github.com/AdhityaRamadhanus/gopatrol/api"
	"github.com/AdhityaRamadhanus/gopatrol/api/handlers"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

func fatalError(err error) {
	log.WithError(err).Error("FATAL ERROR")
	os.Exit(1)
}

func runDaemon(cliContext *cli.Context) {
	godotenv.Load()

	if os.Getenv("ENV") == "Production" {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(log.WarnLevel)
		log.SetOutput(&lumberjack.Logger{
			Filename:   cliContext.String("log"),
			MaxSize:    100,
			MaxAge:     14,
			MaxBackups: 10,
		})
	} else {
		switch cliContext.String("log") {
		case "stdout":
			log.SetOutput(os.Stdout)
		case "stderr":
			log.SetOutput(os.Stderr)
		default:
			log.SetOutput(&lumberjack.Logger{
				Filename:   cliContext.String("log"),
				MaxSize:    100,
				MaxAge:     14,
				MaxBackups: 10,
			})
		}
	}

	intervalRaw := "1m"
	if cliContext.NArg() > 0 {
		intervalRaw = cliContext.Args().Get(0)
	}

	mongoURI := os.Getenv("MONGODB_URI")

	session, err := mgo.Dial(mongoURI)
	defer session.Close()
	if err != nil {
		fatalError(err)
	}

	log.WithFields(log.Fields{
		"mongodb_uri": mongoURI,
	}).Info("connected to mongodb")

	// Creating Service
	loggingService := mongo.NewLoggingService(session, "Logs")
	endpointService := mongo.NewEndpointService(session)

	daemon := &daemon.Daemon{
		Checkup:         &gopatrol.Checkup{},
		LoggingService:  loggingService,
		EndpointService: endpointService,
	}

	// Setting Up Slack Notifier
	slackNotifier := gopatrol.NewSlackNotifier(os.Getenv("SLACK_TOKEN"), os.Getenv("SLACK_CHANNEL"))
	daemon.Notifier = append(daemon.Notifier, slackNotifier)

	// Set Check Interval
	interval, _ := time.ParseDuration(intervalRaw)
	daemon.CheckInterval = interval

	go daemon.Run()

	log.WithFields(log.Fields{
		"interval": interval,
	}).Info("Running daemon service")

	checkersHandler := &handlers.CheckersHandler{
		EndpointService: endpointService,
	}

	api := api.NewApi()
	api.Handlers = append(api.Handlers, checkersHandler)
	api.InitHandler()
	srv := api.CreateServer()
	// Handle SIGINT, SIGTERN, SIGHUP signal from OS
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	go func() {
		<-termChan
		log.Warn("Receiving signal, Shutting down server")
		srv.Close()
	}()

	log.WithField("URL", api.Config.Address).Info("Gopatrol API Server is running")
	fatalError(srv.ListenAndServe())
}
