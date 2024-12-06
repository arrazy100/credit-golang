package models

import (
	"credit/models/base"

	"github.com/google/uuid"
)

type DebtorToUser struct {
	Debtor Debtor    `gorm:"foreignKey:DebtorID"`
	User   base.User `gorm:"foreignKey:UserID"`

	DebtorID uuid.UUID `json:"debtor_id" gorm:"type:uuid;not null"`
	UserID   uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
}
