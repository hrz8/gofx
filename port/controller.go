package port

import "github.com/labstack/echo/v4"

type OrderController interface {
	CreateOrder(c echo.Context) error
}

type UserController interface {
	GetUser(c echo.Context) error
}
