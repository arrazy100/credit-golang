package models

import (
	"credit/models/base"
	"credit/models/enums"
	"math/big"

	"github.com/google/uuid"
)

type DebtorTenorLimit struct {
	base.BaseAuditUser
	Debtor         Debtor               `json:"debtor" gorm:"foreignKey:DebtorID"`
	TenorLimitType enums.TenorLimitType `json:"tenor_limit_type" gorm:"type:int;not null"`
	TenorDuration  int                  `json:"tenor_duration" gorm:"type:int;not null"`
	LimitAmount    big.Float            `json:"limit_amount" gorm:"type:decimal(18,2);not null"`

	DebtorID uuid.UUID `json:"debtor_id" gorm:"type:uuid;not null"`
}
