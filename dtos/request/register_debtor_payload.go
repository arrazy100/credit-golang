package request

type RegisterDebtorPayload struct {
	NIK                string `json:"nik" validate:"required,min=16,max=25"`
	FullName           string `json:"full_name" validate:"required,min=1,max=255"`
	LegalName          string `json:"legal_name" validate:"required,min=1,max=255"`
	PlaceOfBirth       string `json:"place_of_birth" validate:"required,min=1,max=255"`
	DateOfBirth        string `json:"date_of_birth" validate:"required,min=1"`
	Salary             string `json:"salary" validate:"required,number"`
	IdentityPictureUrl string `json:"identity_picture_url" validate:"required,url,min=1,max=2048"`
	SelfiePictureUrl   string `json:"selfie_picture_url" validate:"required,url,min=1,max=2048"`
}
