package cmd

import (
	"fmt"
	"hrz8/gofx/port"
	"time"
)

type httpServer struct {
	ord port.OrderService
}

func (s *httpServer) Run() {
	ticker := time.NewTicker(1500 * time.Millisecond)
	go func() {
		for {
			<-ticker.C
			order := s.ord.CreateOrder("ilham")
			fmt.Println("http: " + order)
		}
	}()
}

func newHttpServer(ord port.OrderService) *httpServer {
	fmt.Println("hello http...")
	srv := &httpServer{ord}

	go srv.Run()

	return srv
}
