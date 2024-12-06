package base

import (
	"credit/models/enums"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID      `gorm:"type:uuid;primary_key"`
	Email    string         `gorm:"type:varchar(255);not null;unique"`
	Password string         `gorm:"type:varchar(255);not null"`
	Role     enums.UserRole `gorm:"type:int;not null"`
}
