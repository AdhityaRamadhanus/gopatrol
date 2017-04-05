package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	checkupgrpc "github.com/AdhityaRamadhanus/checkupd/grpc"
	checkupservice "github.com/AdhityaRamadhanus/checkupd/grpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	portService = flag.String("port", ":9009", "port for grpc server")
	configFile  = flag.String("config", "checkup.json", "config file path")
)

func fatalErr(err error) {
	log.Println(err)
	os.Exit(1)
}

func main() {
	flag.Parse()
	serviceHandler, err := checkupgrpc.NewServiceHandler(*configFile)

	if err != nil {
		fatalErr(err)
	}

	// Handle SIGINT, SIGTERN, SIGHUP signal from OS
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	// Prepare Server
	listener, err := net.Listen("tcp", *portService)
	if err != nil {
		fatalErr(err)
	}
	tcpServer := grpc.NewServer()
	checkupservice.RegisterCheckupServer(tcpServer, serviceHandler)
	reflection.Register(tcpServer)
	// Signal Handler
	go func() {
		<-termChan
		log.Println("Tcp Server is Shutting down")
		if err := serviceHandler.Serialize(); err != nil {
			log.Println("Failed to save checkup.json", err)
		}
		tcpServer.GracefulStop()
	}()

	go serviceHandler.Run()
	// Checkup Goroutine

	log.Println("Tcp server is running at ", *portService)
	if err := tcpServer.Serve(listener); err != nil {
		fatalErr(err)
	}
}
