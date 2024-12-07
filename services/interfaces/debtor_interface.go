package interfaces

import (
	"credit/dtos/request"
	"credit/dtos/response"
	custom_errors "credit/errors"

	"github.com/google/uuid"
)

type DebtorInterface interface {
	RegisterDebtor(userID uuid.UUID, payload request.RegisterDebtorPayload) (*response.RegisterDebtorResponse, int, *custom_errors.ErrorValidation)
	DetailDebtor(userID uuid.UUID) (*response.DebtorResponse, int, *custom_errors.ErrorValidation)
}
