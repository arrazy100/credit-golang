package interfaces

import "credit/dtos/response"

type AuthInterface interface {
	Login(email, password string) (*response.LoginResponse, error)
}
