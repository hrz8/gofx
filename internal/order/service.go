package order

import (
	"fmt"
	"hrz8/gofx/internal/core"
	"hrz8/gofx/model"
	"hrz8/gofx/port"
)

type service struct {
	logger *core.Logger
	config *core.Config
	uSvc   port.UserService
}

func NewService(log *core.Logger, cfg *core.Config, uSvc port.UserService) *service {
	return &service{
		logger: log,
		config: cfg,
		uSvc:   uSvc,
	}
}

func (s *service) PersistOrderData(shouldErr bool) error {
	var user *model.User

	user, err := s.uSvc.GetUserFromDB(false)
	if err != nil {
		return err
	}

	s.logger.Info(fmt.Sprintf("creating order on %s for user %s...", s.config.AppName, user.Name))

	s.logger.Info("persisting order data into database")
	if shouldErr {
		return fmt.Errorf("persist order data error")
	}
	return nil
}
