package port

import (
	"hrz8/gofx/model"
)

type UserService interface {
	GetUser(name string) model.User
}

type OrderService interface {
	CreateOrder(username string) string
}
