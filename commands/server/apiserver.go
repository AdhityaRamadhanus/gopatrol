package daemon

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/AdhityaRamadhanus/gopatrol/mongo"
	log "github.com/Sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"

	"github.com/AdhityaRamadhanus/gopatrol/api"
	"github.com/AdhityaRamadhanus/gopatrol/api/handlers"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

func runApiServer(cliContext *cli.Context) {
	godotenv.Load()
	switch os.Getenv("ENV") {
	case "Production":
		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(log.WarnLevel)
		log.SetOutput(&lumberjack.Logger{
			Filename:   cliContext.String("log"),
			MaxSize:    100,
			MaxAge:     14,
			MaxBackups: 10,
		})
	default:
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
	checkersService := mongo.NewCheckersService(session, "Checkers")
	eventService := mongo.NewEventService(session, "Events")
	// cacheService := cache.NewBigCacheService()
	usersService := mongo.NewUsersService(session, "Users")

	checkersHandler := &handlers.CheckersHandler{
		CheckerService: checkersService,
		// CacheService:   cacheService,
	}

	eventsHandler := &handlers.EventsHandlers{
		EventService: eventService,
		// CacheService: cacheService,
	}

	usersHandler := &handlers.UsersHandler{
		UsersService: usersService,
	}

	address := cliContext.String("address")

	api := api.NewApi()
	api.Handlers = append(api.Handlers, checkersHandler)
	api.Handlers = append(api.Handlers, eventsHandler)
	api.Handlers = append(api.Handlers, usersHandler)
	if cliContext.String("proto") != "unix" {
		api.InitHandler(false)
		srv := api.CreateServer(false)
		// Handle SIGINT, SIGTERN, SIGHUP signal from OS
		termChan := make(chan os.Signal, 1)
		signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		go func() {
			<-termChan
			log.Warn("Receiving signal, Shutting down server")
			srv.Close()
		}()

		log.WithField("URL", address).Info("Gopatrol API Server is running")
		fatalError(srv.ListenAndServe())
	} else {
		api.InitHandler(true)
		srv := api.CreateServer(true)
		// Handle SIGINT, SIGTERN, SIGHUP signal from OS
		termChan := make(chan os.Signal, 1)
		signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		go func() {
			<-termChan
			log.Warn("Receiving signal, Shutting down server")
			srv.Close()
		}()
		unixListener, err := net.Listen("unix", address)
		if err != nil {
			fatalError(err)
		}
		log.WithField("URL", address).Info("Gopatrol API Server is running in unix domain socket")
		fatalError(srv.Serve(unixListener))
	}
}
