package main

import (
	"net/http"

	"github.com/hrz8/gofx"
	"github.com/hrz8/gofx/port"

	"github.com/labstack/echo/v4"
)

func RegisterRouters(oc port.OrderController, uc port.UserController) *gofx.Router {
	router := gofx.NewRouter()

	router.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "pong")
	})
	router.POST("/order", oc.CreateOrder)
	router.GET("/user", uc.GetUser)

	return router
}
