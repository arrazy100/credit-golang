package request

import (
	"credit/models"

	"github.com/google/uuid"
)

type DebtorTransactionPayload struct {
	OTR          string    `json:"otr" validate:"required,number"`
	AssetName    string    `json:"asset_name" validate:"required,min=1,max=255"`
	TenorLimitID uuid.UUID `json:"tenor_limit_id" validate:"required,uuid"`
}

type DebtorTransactionMapResult struct {
	DebtorTransaction *models.DebtorTransaction
	DebtorTenorLimit  *models.DebtorTenorLimit
	DebtorInstallment *models.DebtorInstallment
	Sequence          *models.Sequence
}
