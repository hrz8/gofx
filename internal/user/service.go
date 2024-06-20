package user

import (
	"fmt"

	"github.com/hrz8/gofx"
	"github.com/hrz8/gofx/config"
	"github.com/hrz8/gofx/model"
)

type service struct {
	logger *gofx.Logger
	config *config.Config
}

func NewService(log *gofx.Logger, cfg *config.Config) *service {
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
