package user

import (
	"hrz8/gofx/port"

	"go.uber.org/fx"
)

var Module = fx.Module("user", fx.Provide(
	fx.Annotate(NewService, fx.As(new(port.UserService))),
	fx.Annotate(NewController, fx.As(new(port.UserController))),
))
