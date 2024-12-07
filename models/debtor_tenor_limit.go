package models

import (
	"credit/models/base"
	"credit/models/enums"

	"github.com/google/uuid"
)

type DebtorTenorLimit struct {
	base.BaseAuditUser
	base.BaseUser
	DebtorID       uuid.UUID            `json:"debtor_id" gorm:"type:uuid;not null"`
	TenorLimitType enums.TenorLimitType `json:"tenor_limit_type" gorm:"type:int;not null"`
	TenorDuration  int                  `json:"tenor_duration" gorm:"type:int;not null"`
	TotalLimit     string               `json:"total_limit" gorm:"type:decimal(18,2);not null"`
	CurrentLimit   string               `json:"current_limit" gorm:"type:decimal(18,2);not null"`

	Debtor Debtor `json:"debtor" gorm:"foreignKey:DebtorID"`
}
