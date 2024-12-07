package base

import (
	"time"

	"github.com/google/uuid"
)

type BaseAuditUser struct {
	BaseAudit
	CreatedBy User  `json:"created_by" gorm:"foreignKey:CreatedByID"`
	UpdatedBy *User `json:"updated_by" gorm:"foreignKey:UpdatedByID"`

	CreatedByID uuid.UUID  `json:"created_by_id" gorm:"type:uuid;not null"`
	UpdatedByID *uuid.UUID `json:"updated_by_id" gorm:"type:uuid"`
}

func (b *BaseAuditUser) SetAuditCreated(timeStamp time.Time, createdById uuid.UUID) {
	b.BaseAudit.SetAuditCreated(timeStamp)
	b.CreatedByID = createdById
}

func (b *BaseAuditUser) SetAuditUpdated(timeStamp time.Time, updatedById uuid.UUID) {
	b.BaseAudit.SetAuditUpdated(timeStamp)
	b.UpdatedByID = &updatedById
}
