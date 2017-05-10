package daemon

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	mgo "gopkg.in/mgo.v2"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"

	checkupgrpc "github.com/AdhityaRamadhanus/gopatrol/grpc"
	checkupservice "github.com/AdhityaRamadhanus/gopatrol/grpc/service"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

func runTLSDaemon(cliContext *cli.Context) {
	intervalRaw := "1m"
	if cliContext.NArg() > 0 {
		intervalRaw = cliContext.Args().Get(0)
	}
	connectionString := "mongodb://localhost:27017/gopatrol"
	session, err := mgo.Dial(connectionString)
	if err != nil {
		return
	}

	serviceHandler, err := checkupgrpc.NewServiceHandler(session)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// Set up process log before anything bad happens
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

	// Set Check Interval
	interval, _ := time.ParseDuration(intervalRaw)
	serviceHandler.CheckInterval = interval

	// Handle SIGINT, SIGTERN, SIGHUP signal from OS
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	// Prepare Server
	listener, err := net.Listen("tcp", cliContext.String("address"))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	creds, err := credentials.NewServerTLSFromFile(cliContext.String("cert"), cliContext.String("key"))
	if err != nil {
		log.Fatalf("Failed to generate credentials %v", err)
	}

	tcpServer := grpc.NewServer(
		grpc.Creds(creds),
	)
	checkupservice.RegisterCheckupServer(tcpServer, serviceHandler)
	reflection.Register(tcpServer)
	// Signal Handler
	go func() {
		<-termChan
		log.Println("Tcp Server is Shutting down")
		tcpServer.GracefulStop()
	}()

	go serviceHandler.Run()
	// Checkup Goroutine

	log.Println("Tcp server is running at ", cliContext.String("address"))
	if err := tcpServer.Serve(listener); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func runDaemon(cliContext *cli.Context) {
	intervalRaw := "1m"
	if cliContext.NArg() > 0 {
		intervalRaw = cliContext.Args().Get(0)
	}

	connectionString := "mongodb://localhost:27017/gopatrol"
	session, err := mgo.Dial(connectionString)
	defer session.Close()

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	serviceHandler, err := checkupgrpc.NewServiceHandler(session)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
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

	// Set Check Interval
	interval, _ := time.ParseDuration(intervalRaw)

	serviceHandler.CheckInterval = interval

	// Handle SIGINT, SIGTERN, SIGHUP signal from OS
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	// Prepare Server
	listener, err := net.Listen(cliContext.String("proto"), cliContext.String("address"))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	tcpServer := grpc.NewServer()
	checkupservice.RegisterCheckupServer(tcpServer, serviceHandler)
	reflection.Register(tcpServer)
	// Signal Handler
	go func() {
		<-termChan
		log.Println("Tcp Server is Shutting down")
		tcpServer.GracefulStop()
	}()

	go serviceHandler.Run()
	// Checkup Goroutine

	log.Println("Tcp server is running at ", listener.Addr().String())
	if err := tcpServer.Serve(listener); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
