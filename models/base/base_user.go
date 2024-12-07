package base

import (
	"github.com/google/uuid"
)

type BaseUser struct {
	UserID uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`

	User User `json:"user" gorm:"foreignKey:UserID"`
}

func (b *BaseUser) SetUser(userID uuid.UUID) {
	b.UserID = userID
}
