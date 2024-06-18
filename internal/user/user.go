package user

import (
	"fmt"
	"hrz8/gofx/model"
)

type userSvc struct{}

func (s *userSvc) GetUser(name string) model.User {
	return model.User{
		Name: name + " ganteng",
	}
}

func NewService() *userSvc {
	fmt.Println("initialize user svc...")
	return &userSvc{}
}
