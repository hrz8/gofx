package order

import (
	"fmt"
	"hrz8/gofx/internal/core"
)

type service struct {
	logger *core.Logger
	config *core.Config
}

func NewService(log *core.Logger, cfg *core.Config) *service {
	return &service{
		logger: log,
		config: cfg,
	}
}

func (s *service) PersistOrderData(err bool) error {
	s.logger.Info("persisting order data into database")
	if err {
		return fmt.Errorf("persist order data error")
	}
	return nil
}
