package config

var (
	DatabaseName = "gopatrol"
)

// Config provides the configuration for the API server
type Config struct {
	Logging    bool
	EnableCors bool
	Address    string
}

func GetDefaultConfig() *Config {
	return &Config{
		Logging:    true,
		EnableCors: true,
		Address:    ":3000",
	}
}
