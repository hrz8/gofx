package config

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	AppVersion string
	AppPort    string `env:"APP_PORT,default=1543"`
	AppName    string `env:"APP_NAME,default=goatsapp"`
	LogLevel   string `env:"LOG_LEVEL,default=DEBUG"`
}

func NewConfig() *Config {
	cfg := &Config{AppVersion: Version}
	if err := envconfig.Process(context.Background(), cfg); err != nil {
		panic("Can't load configuration file")
	}

	return cfg
}