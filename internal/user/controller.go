package user

import (
	"net/http"

	"github.com/hrz8/gofx"
	"github.com/hrz8/gofx/config"
	"github.com/hrz8/gofx/port"

	"github.com/labstack/echo/v4"
)

type controller struct {
	logger *gofx.Logger
	config *config.Config
	svc    port.UserService
}

func NewController(log *gofx.Logger, cfg *config.Config, svc port.UserService) *controller {
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
