package models

import (
	"credit/models/base"
	"credit/models/enums"
	"time"

	"github.com/google/uuid"
)

type DebtorInstallmentLine struct {
	base.BaseAuditUser
	base.BaseUser
	DebtorInstallmentID uuid.UUID                         `json:"debtor_installment_id" gorm:"type:uuid;not null;"`
	DueDate             time.Time                         `json:"due_date" gorm:"type:date;not null;"`
	InstallmentNumber   int                               `json:"installment_number" gorm:"type:int;not null;"`
	InstallmentAmount   string                            `json:"installment_amount" gorm:"type:decimal(18,2);not null;"`
	PaymentDate         *time.Time                        `json:"payment_date" gorm:"type:date;"`
	Status              enums.DebtorInstallmentLineStatus `json:"status" gorm:"type:int;not null;"`

	DebtorInstallment DebtorInstallment `json:"debtor_installment" gorm:"foreignKey:DebtorInstallmentID"`
}
