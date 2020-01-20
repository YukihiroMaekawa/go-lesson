package main

import (
	"log"
	"net"

	"github.com/YukihiroMaekawa/go-lesson/internal/grpc/customer"
	"github.com/YukihiroMaekawa/go-lesson/internal/grpc/order"
	"github.com/YukihiroMaekawa/go-lesson/internal/services"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	paymentService := &services.CustomerService{}
	customer.RegisterCustomerServiceServer(s, paymentService)

	orderService := &services.OrderService{}
	order.RegisterOrderServiceServer(s, orderService)

	log.Printf("Listening on %v", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
