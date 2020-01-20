package services

import (
	"fmt"

	"github.com/YukihiroMaekawa/go-lesson/internal/grpc/customer"
	"golang.org/x/net/context"
)

type CustomerService struct{}

func (s *CustomerService) RegistCustomer(ctx context.Context, request *customer.RegistCustomerRequest) (*customer.RegistCustomerResponse, error) {
	fmt.Printf("%+v\n", request)
	resp := customer.RegistCustomerResponse{
		CustomerID: "0001",
	}
	return &resp, nil
}
