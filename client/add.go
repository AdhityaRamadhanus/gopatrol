package main

import (
	"context"
	// "flag"
	"log"

	"github.com/AdhityaRamadhanus/checkup-server/checkupservice"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:9009"
	defaultName = "world"
)

func main() {
	// var blogName = flag.String("name", "", "blog name")
	// var blogType = flag.String("type", "", "blog type")
	// var blogPort = flag.Int("port", 80, "blog port")
	// flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := checkupservice.NewCheckupClient(conn)

	r, err := c.AddHttpEndpoint(context.Background(), &checkupservice.HttpEndpointRequest{
		Endpoint: &checkupservice.GenericEndpointRequest{
			Name:     "Simpan Video",
			Url:      "http://simpanvideo.com",
			Attempts: 5,
		},
	})

	if err != nil {
		log.Fatalf("could not create blog: %v", err)
	}
	log.Printf(r.Message)
}
