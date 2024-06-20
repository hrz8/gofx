package order

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
	svc    port.OrderService
}

func NewController(log *gofx.Logger, cfg *config.Config, svc port.OrderService) *controller {
	return &controller{
		logger: log,
		config: cfg,
		svc:    svc,
	}
}

func (ctrl *controller) CreateOrder(c echo.Context) error {
	if err := ctrl.svc.PersistOrderData(false); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "success create order")
}
