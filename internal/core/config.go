package core

import (
	"context"
	"fmt"

	"github.com/sethvargo/go-envconfig"
)

type AppVersion string

type Config struct {
	AppVersion AppVersion
	AppPort    string `env:"APP_PORT,default=1543"`
	AppName    string `env:"APP_NAME,default=goatsapp"`
	LogLevel   string `env:"LOG_LEVEL,default=DEBUG"`
}

func NewConfig(version AppVersion) *Config {
	fmt.Println("load version: " + version)
	cfg := &Config{AppVersion: version}
	if err := envconfig.Process(context.Background(), cfg); err != nil {
		panic("Can't load configuration file")
	}

	return cfg
}
