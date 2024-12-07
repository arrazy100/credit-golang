package response

import (
	"credit/models/enums"
)

type RegisterDebtorResponse struct {
	NIK                string                     `json:"nik"`
	FullName           string                     `json:"full_name"`
	LegalName          string                     `json:"legal_name"`
	PlaceOfBirth       string                     `json:"place_of_birth"`
	DateOfBirth        string                     `json:"date_of_birth"`
	Salary             string                     `json:"salary"`
	IdentityPictureUrl string                     `json:"identity_picture_url"`
	SelfiePictureUrl   string                     `json:"selfie_picture_url"`
	TenorLimits        []DebtorTenorLimitResponse `json:"debtor_tenor_limits"`
}

type DebtorTenorLimitResponse struct {
	TenorLimitType enums.TenorLimitType `json:"tenor_limit_type"`
	TenorDuration  int                  `json:"tenor_duration"`
	TotalLimit     string               `json:"total_limit"`
	CurrentLimit   string               `json:"current_limit"`
}
