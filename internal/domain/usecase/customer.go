//go:generate mockgen -destination mocks/mock_customer.go -package=mocks -source=customer.go
package usecase

import (
	"context"
	"fmt"
	"github.com/YukihiroMaekawa/go-lesson/internal/domain/entity"
)

type dber interface {
	CreateCustomer(request *entity.RegisterCustomerRequest) (string, error)
}

type Customer struct {
	DB dber
}

func (s *Customer) CreateCustomer(ctx context.Context, request *entity.RegisterCustomerRequest) (string, error) {
	cus, err := s.DB.CreateCustomer(request)
	if err != nil {
		return "ERR", err
	}
	return fmt.Sprintf("%s:%s", cus, request.Name), nil
}
