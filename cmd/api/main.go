package main

import (
	"log"
	"os"

	"github.com/AdhityaRamadhanus/gopatrol/api"
	"github.com/AdhityaRamadhanus/gopatrol/api/handlers"
	"github.com/AdhityaRamadhanus/gopatrol/mongo"
	"github.com/joho/godotenv"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	session, err := mgo.Dial(os.Getenv("MONGODB_URI"))
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	// redisClient := redis.NewClient(&redis.Options{
	// 	Addr: os.Getenv("REDIS_URI"),
	// 	DB:   0})
	// defer redisClient.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Product Handler
	// productHandler := &handlers.ProductHandler{
	// 	PrService: mongo.NewProductService(session),
	// 	// CacheService:   cache.NewCacheService(redisClient),
	// }
	endpointHandler := &handlers.EndpointHandler{
		EndpointService: mongo.NewEndpointService(session),
	}
	apiServer := api.NewApiServer()
	apiServer.AddHandler(endpointHandler)
	apiServer.InitHandler()
	log.Println("Gopatrol api server running at", apiServer.Config.Address)
	log.Fatal(apiServer.ListenAndServe())
}
