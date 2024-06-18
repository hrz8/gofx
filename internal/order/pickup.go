package order

import (
	"fmt"
	"hrz8/gofx/port"
)

type pickupOrderSvc struct {
	uSvc port.UserService
}

func (s *pickupOrderSvc) CreateOrder(username string) string {
	user := s.uSvc.GetUser(username)
	return "self-pickup order for " + user.Name + " created"
}

func NewPickupService(uSvc port.UserService) *pickupOrderSvc {
	fmt.Println("initialize self-pickup order svc...")
	return &pickupOrderSvc{uSvc}
}
