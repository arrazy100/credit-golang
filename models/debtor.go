package models

import (
	"credit/models/base"
	"math/big"
	"time"
)

type Debtor struct {
	base.BaseAuditUser
	NIK                string    `json:"nik" gorm:"varchar(25);not null;unique"`
	FullName           string    `json:"full_name" gorm:"varchar(255);not null"`
	LegalName          string    `json:"legal_name" gorm:"varchar(255);not null"`
	PlaceOfBirth       string    `json:"place_of_birth" gorm:"varchar(255);not null"`
	DateOfBirth        time.Time `json:"date_of_birth" gorm:"not null"`
	Salary             big.Float `json:"salary" gorm:"type:decimal(18,2);not null"`
	IdentityPictureUrl string    `json:"identity_picture_url" gorm:"varchar(2048);not null"`
	SelfiePictureUrl   string    `json:"selfie_picture_url" gorm:"varchar(2048);not null"`
}
