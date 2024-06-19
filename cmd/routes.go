package cmd

import (
	"hrz8/gofx/internal/core"
	"hrz8/gofx/port"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterRouters(oc port.OrderController, uc port.UserController) *core.Router {
	router := core.NewRouter()

	router.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "pong")
	})
	router.POST("/order", oc.CreateOrder)
	router.GET("/user", uc.GetUser)

	return router
}
