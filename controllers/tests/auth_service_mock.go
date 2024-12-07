package controllers_tests

import (
	"credit/dtos/request"
	"credit/dtos/response"
	custom_errors "credit/errors"

	"github.com/stretchr/testify/mock"
)

type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) Login(payload request.LoginPayload) (*response.LoginResponse, int, *custom_errors.ErrorValidation) {
	args := m.Called(payload)
	if args.Get(0) == nil {
		return nil, args.Get(1).(int), args.Get(2).(*custom_errors.ErrorValidation)
	}
	return args.Get(0).(*response.LoginResponse), args.Get(1).(int), nil
}

func (m *MockAuthService) RegisterUser(payload request.RegisterPayload) (*response.RegisterResponse, int, *custom_errors.ErrorValidation) {
	args := m.Called(payload)
	if args.Get(0) == nil {
		return nil, args.Get(1).(int), args.Get(2).(*custom_errors.ErrorValidation)
	}
	return args.Get(0).(*response.RegisterResponse), args.Get(1).(int), nil
}
