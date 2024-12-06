package models

import (
	"credit/models/base"
	"math/big"
	"time"

	"github.com/google/uuid"
)

type DebtorInstallment struct {
	base.BaseAuditUser
	DebtorTransaction  DebtorTransaction `json:"debtor_transaction" gorm:"foreignKey:DebtorTransactionID"`
	Debtor             Debtor            `json:"debtor" gorm:"foreignKey:DebtorID"`
	TenorLimitType     DebtorTenorLimit  `json:"tenor_limit_type" gorm:"type:int;not null"`
	TenorDuration      int               `json:"tenor_duration" gorm:"type:int;not null"`
	MonthlyInstallment big.Float         `json:"monthly_installment" gorm:"type:decimal(18,2);not null"`
	TotalInstallment   big.Float         `json:"total_installment" gorm:"type:decimal(10,2);not null"`
	StartDatePeriod    time.Time         `json:"start_date_period" gorm:"type:date;not null"`
	EndDatePeriod      time.Time         `json:"end_date_period" gorm:"type:date;not null"`

	DebtorTransactionID uuid.UUID `json:"debtor_transaction_id" gorm:"type:uuid;not null"`
	DebtorID            uuid.UUID `json:"debtor_id" gorm:"type:uuid;not null"`
}
