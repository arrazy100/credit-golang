package models

import (
	"credit/models/base"
	"time"

	"github.com/google/uuid"
)

type DebtorInstallment struct {
	base.BaseAuditUser
	base.BaseUser
	DebtorTransactionID uuid.UUID        `json:"debtor_transaction_id" gorm:"type:uuid;not null"`
	TenorLimitType      DebtorTenorLimit `json:"tenor_limit_type" gorm:"type:int;not null"`
	TenorDuration       int              `json:"tenor_duration" gorm:"type:int;not null"`
	MonthlyInstallment  string           `json:"monthly_installment" gorm:"type:decimal(18,2);not null"`
	TotalInstallment    string           `json:"total_installment" gorm:"type:decimal(10,2);not null"`
	StartDatePeriod     time.Time        `json:"start_date_period" gorm:"type:date;not null"`
	EndDatePeriod       time.Time        `json:"end_date_period" gorm:"type:date;not null"`

	DebtorTransaction DebtorTransaction `json:"debtor_transaction" gorm:"foreignKey:DebtorTransactionID"`
}
