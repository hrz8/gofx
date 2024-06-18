package cmd

import (
	"hrz8/gofx/internal/order"
	"hrz8/gofx/internal/user"
	"hrz8/gofx/port"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

const (
	ORDER_TYPE = "pickup"
)

var AppCmd = &cobra.Command{
	Use: "serve",
	Run: run,
}

func loadProviders() fx.Option {
	var constructors []interface{}
	var orderSvc interface{}

	if ORDER_TYPE == "delivery" {
		orderSvc = order.NewDeliveryService
	} else {
		orderSvc = order.NewPickupService
	}

	constructors = append(constructors, fx.Annotate(user.NewService, fx.As(new(port.UserService))))
	constructors = append(constructors, fx.Annotate(orderSvc, fx.As(new(port.OrderService))))

	return fx.Provide(constructors...)
}

func loadServers() fx.Option {
	return fx.Provide(newHttpServer, newGrpcServer)
}

func run(_ *cobra.Command, _ []string) {
	fx.New(
		loadProviders(),
		loadServers(),
		fx.Invoke(func(*httpServer) {}, func(*grpcServer) {}),
	).Run()
}
