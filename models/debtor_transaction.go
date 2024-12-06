package models

import (
	"credit/models/base"
	"credit/models/enums"
	"math/big"

	"github.com/google/uuid"
)

type DebtorTransaction struct {
	base.BaseAuditUser
	Debtor         Debtor                        `json:"debtor" gorm:"foreignKey:DebtorID"`
	ContractNumber string                        `json:"contract_number" gorm:"type:varchar(255);not null"`
	OTR            big.Float                     `json:"otr" gorm:"type:decimal(18,2);not null"`
	AdminFee       big.Float                     `json:"admin_fee" gorm:"type:decimal(18,2);not null"`
	TotalLoan      big.Float                     `json:"total_loan" gorm:"type:decimal(18,2);not null"`
	TotalInterest  big.Float                     `json:"total_interest" gorm:"type:decimal(18,2);not null"`
	AssetName      string                        `json:"asset_name" gorm:"type:varchar(255);not null"`
	Status         enums.DebtorTransactionStatus `json:"status" gorm:"type:int;not null"`

	DebtorID uuid.UUID `json:"debtor_id" gorm:"type:uuid;not null"`
}
