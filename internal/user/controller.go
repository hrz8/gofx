package user

import (
	"hrz8/gofx/config"
	"hrz8/gofx/core"
	"hrz8/gofx/port"
	"net/http"

	"github.com/labstack/echo/v4"
)

type controller struct {
	logger *core.Logger
	config *config.Config
	svc    port.UserService
}

func NewController(log *core.Logger, cfg *config.Config, svc port.UserService) *controller {
	return &controller{
		logger: log,
		config: cfg,
		svc:    svc,
	}
}

func (ctrl *controller) GetUser(c echo.Context) error {
	user, err := ctrl.svc.GetUserFromDB(false)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, user)
}
