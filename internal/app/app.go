package app

import (
	"database/sql"
	"log"
	"net"

	"github.com/YukihiroMaekawa/go-lesson/internal/domain/usecase"
	"github.com/YukihiroMaekawa/go-lesson/internal/gateway/rdb"
	"github.com/YukihiroMaekawa/go-lesson/internal/gateway/service"
	"github.com/YukihiroMaekawa/go-lesson/internal/grpc/customer"
	"github.com/YukihiroMaekawa/go-lesson/internal/grpc/order"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

func Start() {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=admin password=admin dbname=pg-go-lesson sslmode=disable")
	if err != nil {
		panic("db error")
	}
	defer db.Close()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	paymentService := service.NewCustomer(&usecase.Customer{DB: rdb.NewCustomer(db)})
	customer.RegisterCustomerServiceServer(s, paymentService)

	orderService := &service.OrderService{}
	order.RegisterOrderServiceServer(s, orderService)

	log.Printf("Listening on %v", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
