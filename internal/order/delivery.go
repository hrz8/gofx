package order

import (
	"fmt"
	"hrz8/gofx/port"
)

type deliveryOrderSvc struct {
	uSvc port.UserService
}

func (s *deliveryOrderSvc) CreateOrder(username string) string {
	user := s.uSvc.GetUser(username)
	return "delivery order for " + user.Name + " created"
}

func NewDeliveryService(uSvc port.UserService) *deliveryOrderSvc {
	fmt.Println("initialize delivery order svc...")
	return &deliveryOrderSvc{uSvc}
}
