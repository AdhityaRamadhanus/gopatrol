package checkupd

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	checkupgrpc "github.com/AdhityaRamadhanus/checkupd/grpc"
	checkupservice "github.com/AdhityaRamadhanus/checkupd/grpc/service"
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
	serviceHandler, err := checkupgrpc.NewServiceHandler(cliContext.String("config"))

	if err != nil {
		log.Println(err)
		os.Exit(1)
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
		if err := serviceHandler.SerializeJSON(); err != nil {
			log.Println("Failed to save checkup.json", err)
		}
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
	serviceHandler, err := checkupgrpc.NewServiceHandler(cliContext.String("config"))

	if err != nil {
		log.Println(err)
		os.Exit(1)
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
		if err := serviceHandler.SerializeJSON(); err != nil {
			log.Println("Failed to save checkup.json", err)
		}
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
