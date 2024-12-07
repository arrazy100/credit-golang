package interfaces

import (
	"credit/dtos/response"
	custom_errors "credit/errors"
)

type AdminInterface interface {
	ListDebtor() (*response.ListDebtorResponse, int, *custom_errors.ErrorValidation)
}
