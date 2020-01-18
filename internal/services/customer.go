package services

import (
	"github.com/YukihiroMaekawa/go-lesson/pkg/customer"
	"golang.org/x/net/context"
)

type customerService struct{}

func (s *customerService) RegistCustomer(ctx context.Context, request *customer.RegistCustomerRequest) (*customer.RegistCustomerResponse, error) {
	return nil, nil
}
