package interfaces

import (
	"credit/dtos/request"
	"credit/dtos/response"
	validations "credit/validations"
	"sync"

	"github.com/google/uuid"
)

type DebtorInterface interface {
	Register(userID uuid.UUID, payload request.RegisterDebtorPayload) (*response.RegisterDebtorResponse, int, *validations.ErrorValidation)
	Detail(userID uuid.UUID) (*response.DebtorResponse, int, *validations.ErrorValidation)
	CreateTransaction(userID uuid.UUID, payload request.DebtorTransactionPayload) (*response.DebtorTransactionResponse, int, *validations.ErrorValidation)
	ListInstallment(userID uuid.UUID) (*response.ListDebtorInstallmentResponse, int, *validations.ErrorValidation)
	PayInstallmentLine(userID uuid.UUID, payload request.DebtorPayInstallmentLinePayload) (*response.DebtorInstallmentLineResponse, int, *validations.ErrorValidation)
	BatchUpdateOverdueInstallmentLine(wg *sync.WaitGroup, errCh chan<- error)
}
