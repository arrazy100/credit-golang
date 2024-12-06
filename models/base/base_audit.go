package base

import (
	"time"

	"github.com/google/uuid"
)

type BaseAudit struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`
}

func (b *BaseAudit) SetAuditCreated(timeStamp time.Time) {
	b.CreatedAt = timeStamp
	b.ID = uuid.New()
}

func (b *BaseAudit) SetAuditUpdated(timeStamp time.Time) {
	b.UpdatedAt = timeStamp
}
