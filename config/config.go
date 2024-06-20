package config

type Config struct {
	AppVersion string
	AppPort    string
	AppName    string
	LogLevel   string
}

func NewConfig() *Config {
	cfg := &Config{
		AppVersion: Version,
		AppPort:    "1543",
		AppName:    "app",
		LogLevel:   "DEBUG",
	}

	return cfg
}
