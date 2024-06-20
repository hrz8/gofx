package main

import (
	"net/http"

	"github.com/hrz8/gofx"
	"github.com/hrz8/gofx/config"
	"github.com/hrz8/gofx/internal/order"
	"github.com/hrz8/gofx/internal/user"
)

func main() {
	cfg := config.NewConfig()

	app := gofx.NewApp(&gofx.Config{
		Version:  cfg.AppVersion,
		Addr:     ":" + cfg.AppPort,
		LogLevel: cfg.LogLevel,
	})
	app.AddModules(order.Module, user.Module)
	app.AddProviders(RegisterRouters, config.NewConfig)
	app.AddServers(gofx.NewHTTPServer)
	app.AddInvokers(func(*http.Server) {})

	app.Start()
}
