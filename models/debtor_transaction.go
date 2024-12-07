package models

import (
	"credit/models/base"
	"credit/models/enums"
)

type DebtorTransaction struct {
	base.BaseAuditUser
	base.BaseUser
	ContractNumber string                        `json:"contract_number" gorm:"type:varchar(255);not null"`
	OTR            string                        `json:"otr" gorm:"type:decimal(18,2);not null"`
	AdminFee       string                        `json:"admin_fee" gorm:"type:decimal(18,2);not null"`
	TotalLoan      string                        `json:"total_loan" gorm:"type:decimal(18,2);not null"`
	TotalInterest  string                        `json:"total_interest" gorm:"type:decimal(18,2);not null"`
	AssetName      string                        `json:"asset_name" gorm:"type:varchar(255);not null"`
	Status         enums.DebtorTransactionStatus `json:"status" gorm:"type:int;not null"`
}
