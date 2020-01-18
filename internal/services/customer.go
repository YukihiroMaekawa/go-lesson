package services

import (
	"github.com/YukihiroMaekawa/go-lesson/pkg/customer"
	"golang.org/x/net/context"
)

type CustomerService struct{}

func (s *CustomerService) RegistCustomer(ctx context.Context, request *customer.RegistCustomerRequest) (*customer.RegistCustomerResponse, error) {
	resp := customer.RegistCustomerResponse{
		CustomerID: "0001",
	}
	return &resp, nil
}
