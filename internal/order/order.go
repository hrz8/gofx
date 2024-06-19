package order

import (
	"hrz8/gofx/port"

	"go.uber.org/fx"
)

var Module = fx.Module("order", fx.Provide(
	fx.Annotate(NewService, fx.As(new(port.OrderService))),
	fx.Annotate(NewController, fx.As(new(port.OrderController))),
))
