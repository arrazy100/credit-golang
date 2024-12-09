package interfaces

import (
	"credit/dtos/response"
	validations "credit/validations"
)

type AdminInterface interface {
	ListDebtor() (*response.ListDebtorResponse, int, *validations.ErrorValidation)
}
