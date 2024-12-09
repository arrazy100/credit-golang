package models

import (
	"credit/models/base"
	"time"

	"github.com/google/uuid"
)

type DebtorInstallment struct {
	base.BaseAuditUser
	base.BaseUser
	DebtorTransactionID uuid.UUID `json:"debtor_transaction_id" gorm:"type:uuid;not null"`
	TenorLimitID        uuid.UUID `json:"tenor_limit_id" gorm:"type:int;not null"`
	MonthlyInstallment  string    `json:"monthly_installment" gorm:"type:decimal(18,2);not null"`
	TotalInstallment    string    `json:"total_installment" gorm:"type:decimal(10,2);not null"`
	StartDatePeriod     time.Time `json:"start_date_period" gorm:"type:date;not null"`
	EndDatePeriod       time.Time `json:"end_date_period" gorm:"type:date;not null"`

	DebtorTransaction DebtorTransaction       `json:"debtor_transaction" gorm:"foreignKey:DebtorTransactionID"`
	DebtorTenorLimit  DebtorTenorLimit        `json:"debtor_tenor_limit" gorm:"foreignKey:TenorLimitID"`
	InstallmentLines  []DebtorInstallmentLine `json:"installment_lines" gorm:"foreignKey:DebtorInstallmentID"`
}
