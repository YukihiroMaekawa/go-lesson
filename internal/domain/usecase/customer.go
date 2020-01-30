package usecase

import (
	"context"

	"github.com/YukihiroMaekawa/go-lesson/internal/domain/entity"
)

type dber interface {
	InsertCustomer(request *entity.RegisterCustomerRequest) (string, error)
}

type Customer struct {
	DB dber
}

func (s *Customer) RegistCustomer(ctx context.Context, request *entity.RegisterCustomerRequest) (string, error) {
	/*
		ロジック。。。
	*/
	return s.DB.InsertCustomer(request)
}
