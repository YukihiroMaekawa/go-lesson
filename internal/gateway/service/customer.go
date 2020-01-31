package service

import (
	"github.com/YukihiroMaekawa/go-lesson/internal/domain/entity"
	"github.com/YukihiroMaekawa/go-lesson/internal/domain/usecase"
	"github.com/YukihiroMaekawa/go-lesson/internal/grpc/customer"
	"golang.org/x/net/context"
)

type Customer struct {
	customer *usecase.Customer
}

func NewCustomer(p *usecase.Customer) *Customer {
	return &Customer{customer: p}
}

func (s *Customer) CreateCustomer(ctx context.Context, request *customer.CreateCustomerRequest) (*customer.CreateCustomerResponse, error) {
	req := &entity.RegisterCustomerRequest{
		Name:     request.GetName(),
		Address1: request.GetAddress1(),
		Address2: request.GetAddress2(),
		Address3: request.GetAddress3(),
	}

	customerNo, err := s.customer.CreateCustomer(ctx, req)
	if err != nil {
		return nil, err
	}

	resp := &customer.CreateCustomerResponse{
		CustomerID: customerNo,
	}
	return resp, nil
}
