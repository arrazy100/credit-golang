package request

import "github.com/google/uuid"

type DebtorPayInstallmentLinePayload struct {
	InstallmentLineID uuid.UUID `json:"installment_line_id" validate:"required,uuid"`
}
