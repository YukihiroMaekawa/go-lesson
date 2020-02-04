//go:generate mockgen -destination mocks/mock_customer.go -package=mocks -source=customer.go
package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/YukihiroMaekawa/go-lesson/internal/domain/entity"
	"github.com/YukihiroMaekawa/go-lesson/internal/domain/usecase/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCustomer_CreateCustomer(t *testing.T) {
	type args struct {
		request *entity.RegisterCustomerRequest
	}
	tests := []struct {
		name          string
		args          args
		expected      string
		expectedError assert.ErrorAssertionFunc
		runMock       func(*mocks.Mockdber, *entity.RegisterCustomerRequest)
	}{
		{
			name: "test1",
			args: args{
				request: &entity.RegisterCustomerRequest{Name: "name1", Address1: "address1", Address2: "address2", Address3: "address3"},
			},
			expected:      "customer:name1",
			expectedError: assert.NoError,
			runMock: func(m *mocks.Mockdber, req *entity.RegisterCustomerRequest) {
				gomock.InOrder(
					m.EXPECT().CreateCustomer(req).
						Return("customer", nil),
				)
			},
		},
		{
			name: "test2",
			args: args{
				request: &entity.RegisterCustomerRequest{Name: "name2", Address1: "address1", Address2: "address2", Address3: "address3"},
			},
			expected:      "customer:name2",
			expectedError: assert.NoError,
			runMock: func(m *mocks.Mockdber, req *entity.RegisterCustomerRequest) {
				gomock.InOrder(
					m.EXPECT().CreateCustomer(req).
						Return("customer", nil),
				)
			},
		},
		{
			name: "test3",
			args: args{
				request: &entity.RegisterCustomerRequest{Name: "ERR", Address1: "address1", Address2: "address2", Address3: "address3"},
			},
			expected:      "ERR",
			expectedError: assert.Error,
			runMock: func(m *mocks.Mockdber, req *entity.RegisterCustomerRequest) {
				gomock.InOrder(
					m.EXPECT().CreateCustomer(req).
						Return("", errors.New("ERR TEST")),
				)
			},
		},
	}

	testCtx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mocks.NewMockdber(ctrl)

			tt.runMock(mock, tt.args.request)

			app := Customer{DB: mock}
			actual, err := app.CreateCustomer(testCtx, tt.args.request)

			tt.expectedError(t, err)
			// assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}

	// type fields struct {
	// 	DB dber
	// }
	// type args struct {
	// 	ctx     context.Context
	// 	request *entity.RegisterCustomerRequest
	// }
	// tests := []struct {
	// 	name      string
	// 	fields    fields
	// 	args      args
	// 	want      string
	// 	assertion assert.ErrorAssertionFunc
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		s := &Customer{
	// 			DB: tt.fields.DB,
	// 		}
	// 		got, err := s.CreateCustomer(tt.args.ctx, tt.args.request)
	// 		tt.assertion(t, err)
	// 		assert.Equal(t, tt.want, got)
	// 	})
	// }
}
