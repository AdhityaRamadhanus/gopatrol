package config

import (
	"os"
)

var (
	DatabaseName = "gopatrol"
	JwtSecret    = os.Getenv("JWT_SECRET")
)

// Config provides the configuration for the API server
type Config struct {
	EnableCors bool
	Address    string
}

func GetDefaultConfig() *Config {
	return &Config{
		EnableCors: true,
		Address:    ":3000",
	}
}
