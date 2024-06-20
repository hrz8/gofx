package cmd

import (
	"net/http"

	"github.com/hrz8/gofx"
	"github.com/hrz8/gofx/config"
	"github.com/hrz8/gofx/internal/order"
	"github.com/hrz8/gofx/internal/user"

	"github.com/spf13/cobra"
)

var AppCmd = &cobra.Command{
	Use: "serve",
	Run: run,
}

func run(_ *cobra.Command, _ []string) {
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
