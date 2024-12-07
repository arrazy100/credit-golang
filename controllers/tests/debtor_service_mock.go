package controllers_tests

import (
	"credit/dtos/request"
	"credit/dtos/response"
	custom_errors "credit/errors"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockDebtorService struct {
	mock.Mock
}

func (m *MockDebtorService) RegisterDebtor(userID uuid.UUID, payload request.RegisterDebtorPayload) (*response.RegisterDebtorResponse, int, *custom_errors.ErrorValidation) {
	args := m.Called(userID, payload)
	if args.Get(0) == nil {
		return nil, args.Get(1).(int), args.Get(2).(*custom_errors.ErrorValidation)
	}
	return args.Get(0).(*response.RegisterDebtorResponse), args.Get(1).(int), nil
}

func (m *MockDebtorService) DetailDebtor(userID uuid.UUID) (*response.DebtorResponse, int, *custom_errors.ErrorValidation) {
	args := m.Called(userID)
	if args.Get(0) == nil {
		return nil, args.Get(1).(int), args.Get(2).(*custom_errors.ErrorValidation)
	}
	return args.Get(0).(*response.DebtorResponse), args.Get(1).(int), nil
}
