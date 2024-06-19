package user

import (
	"fmt"
	"hrz8/gofx/config"
	"hrz8/gofx/internal/core"
	"hrz8/gofx/model"
)

type service struct {
	logger *core.Logger
	config *config.Config
}

func NewService(log *core.Logger, cfg *config.Config) *service {
	return &service{
		logger: log,
		config: cfg,
	}
}

func (s *service) GetUserFromDB(err bool) (*model.User, error) {
	s.logger.Info("attempt to fetch user from database")
	if err {
		return nil, fmt.Errorf("fetch user error")
	}
	return &model.User{Name: "yasa"}, nil
}
