package cmd

import (
	"fmt"
	"hrz8/gofx/port"
	"time"
)

type grpcServer struct {
	ord port.OrderService
}

func (s *grpcServer) Run() {
	ticker := time.NewTicker(1200 * time.Millisecond)
	go func() {
		for {
			<-ticker.C
			order := s.ord.CreateOrder("ahmad")
			fmt.Println("grpc: " + order)
		}
	}()
}

func newGrpcServer(ord port.OrderService) *grpcServer {
	fmt.Println("hello grpc...")
	srv := &grpcServer{ord}

	go srv.Run()

	return srv
}
