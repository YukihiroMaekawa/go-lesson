//go:generate mockgen -destination mocks/mock_customer.go -package=mocks -source=customer.go
package usecase

import (
	"context"

	"github.com/YukihiroMaekawa/go-lesson/internal/domain/entity"
)

type dber interface {
	CreateCustomer(request *entity.RegisterCustomerRequest) (string, error)
}

type Customer struct {
	DB dber
}

func (s *Customer) CreateCustomer(ctx context.Context, request *entity.RegisterCustomerRequest) (string, error) {
	/*
		ロジック。。。
	*/
	return s.DB.CreateCustomer(request)
}
