package port

import "hrz8/gofx/model"

type OrderService interface {
	PersistOrderData(err bool) error
}

type UserService interface {
	GetUserFromDB(err bool) (*model.User, error)
}
