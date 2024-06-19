package order

import (
	"fmt"
	"hrz8/gofx/internal/core"
	"hrz8/gofx/model"
	"hrz8/gofx/port"
	"net/http"

	"github.com/labstack/echo/v4"
)

type controller struct {
	logger *core.Logger
	config *core.Config
	svc    port.OrderService
	uSvc   port.UserService
}

func NewController(log *core.Logger, cfg *core.Config, svc port.OrderService, uSvc port.UserService) *controller {
	return &controller{
		logger: log,
		config: cfg,
		svc:    svc,
		uSvc:   uSvc,
	}
}

func (ctrl *controller) CreateOrder(c echo.Context) error {
	var err error
	var user *model.User

	user, err = ctrl.uSvc.GetUserFromDB(false)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	ctrl.logger.Info(fmt.Sprintf("creating order on %s for user %s...", ctrl.config.AppName, user.Name))

	if err = ctrl.svc.PersistOrderData(false); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "success create order")
}
