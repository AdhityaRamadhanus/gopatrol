package config

var (
	DatabaseName = "gopatrol"
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
