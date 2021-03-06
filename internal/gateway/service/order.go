package service

import (
	"fmt"

	"github.com/YukihiroMaekawa/go-lesson/internal/grpc/order"
	"golang.org/x/net/context"
)

type OrderService struct{}

func (s *OrderService) RegistOrder(ctx context.Context, request *order.RegistOrderRequest) (*order.RegistOrderResponse, error) {
	fmt.Printf("%+v\n", request)
	resp := order.RegistOrderResponse{
		OrderID: "order-1",
	}
	return &resp, nil
}
