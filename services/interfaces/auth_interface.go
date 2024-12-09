package interfaces

import (
	"credit/dtos/request"
	"credit/dtos/response"
	validations "credit/validations"
)

type AuthInterface interface {
	Login(payload request.LoginPayload) (*response.LoginResponse, int, *validations.ErrorValidation)
	RegisterUser(payload request.RegisterPayload) (*response.RegisterResponse, int, *validations.ErrorValidation)
}
