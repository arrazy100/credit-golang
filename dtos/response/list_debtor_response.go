package response

type ListDebtorResponse struct {
	Data  []DebtorResponse `json:"data"`
	Total int              `json:"total"`
}

type DebtorResponse struct {
	ID                 string                     `json:"id"`
	UserID             string                     `json:"user_id"`
	NIK                string                     `json:"nik"`
	FullName           string                     `json:"full_name"`
	LegalName          string                     `json:"legal_name"`
	PlaceOfBirth       string                     `json:"place_of_birth"`
	DateOfBirth        string                     `json:"date_of_birth"`
	Salary             string                     `json:"salary"`
	IdentityPictureUrl string                     `json:"identity_picture_url"`
	SelfiePictureUrl   string                     `json:"selfie_picture_url"`
	TenorLimits        []DebtorTenorLimitResponse `json:"tenor_limits"`
}
