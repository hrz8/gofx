package cmd

import (
	"hrz8/gofx/internal/core"
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
	app := core.NewApp(Version)
	app.AddModules(
		order.Module,
		user.Module,
	)
	app.AddProviders(RegisterRouters)
	app.AddServers(core.NewHTTPServer)
	app.AddInvokers(func(*http.Server) {})

	app.Start()
}
