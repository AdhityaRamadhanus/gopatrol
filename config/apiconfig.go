package config

import (
	"os"
)

var (
	DatabaseName = "gopatrol"
	JwtSecret    = os.Getenv("JWT_SECRET")
	MongoURI     = os.Getenv("MONGODB_URI")
)

func SetTestingConfig() {
	DatabaseName = "gopatrol-test"
	JwtSecret = "gopatrol-test-jwt-secret"
	MongoURI = "mongodb://localhost:27017"
}
