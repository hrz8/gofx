package cmd

import (
	"hrz8/gofx/config"
	"hrz8/gofx/core"
	"hrz8/gofx/internal/order"
	"hrz8/gofx/internal/user"
	"net/http"

	"github.com/spf13/cobra"
)

var AppCmd = &cobra.Command{
	Use: "serve",
	Run: run,
}

func run(_ *cobra.Command, _ []string) {
	cfg := config.NewConfig()

	app := core.NewApp(&core.AppConfig{
		Version:  cfg.AppVersion,
		Addr:     ":" + cfg.AppPort,
		LogLevel: cfg.LogLevel,
	})
	app.AddModules(order.Module, user.Module)
	app.AddProviders(RegisterRouters, config.NewConfig)
	app.AddServers(core.NewHTTPServer)
	app.AddInvokers(func(*http.Server) {})

	app.Start()
}
