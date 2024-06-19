package order

import (
	"hrz8/gofx/config"
	"hrz8/gofx/internal/core"
	"hrz8/gofx/port"
	"net/http"

	"github.com/labstack/echo/v4"
)

type controller struct {
	logger *core.Logger
	config *config.Config
	svc    port.OrderService
}

func NewController(log *core.Logger, cfg *config.Config, svc port.OrderService) *controller {
	return &controller{
		logger: log,
		config: cfg,
		svc:    svc,
	}
}

func (ctrl *controller) CreateOrder(c echo.Context) error {
	// time.Sleep(15 * time.Second)
	if err := ctrl.svc.PersistOrderData(false); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "success create order")
}
