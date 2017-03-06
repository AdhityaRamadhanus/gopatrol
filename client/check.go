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

	r, err := c.ListEndpoint(context.Background(), &checkupservice.ListEndpointRequest{
		Check: true,
	})

	if err != nil {
		log.Fatalf("could not create blog: %v", err)
	}
	for _, endpoint := range r.Endpoints {
		log.Println(endpoint.Name, " ", endpoint.Url, " ", endpoint.Status)
	}
}
