package interfaces

import (
	"credit/dtos/request"
	"credit/dtos/response"
	custom_errors "credit/errors"
)

type AuthInterface interface {
	Login(payload request.LoginPayload) (*response.LoginResponse, int, *custom_errors.ErrorValidation)
	RegisterUser(payload request.RegisterPayload) (*response.RegisterResponse, int, *custom_errors.ErrorValidation)
}
