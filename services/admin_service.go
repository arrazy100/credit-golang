package services

import (
	"credit/dtos/response"
	custom_errors "credit/errors"
	"credit/models"
	"credit/services/interfaces"
	"credit/utils"

	"gorm.io/gorm"
)

var _ interfaces.AdminInterface = (*AdminService)(nil)

const ()

type AdminService struct {
	db *gorm.DB
}

func NewAdminService(db *gorm.DB) *AdminService {
	return &AdminService{db: db}
}

func (s *AdminService) ListDebtor() (*response.ListDebtorResponse, int, *custom_errors.ErrorValidation) {
	var debtors []models.Debtor
	if err := s.db.Find(&debtors).Error; err != nil {
		return nil, statusServerError, custom_errors.Convert(err)
	}

	return mapDebtorsAsResponse(debtors), statusSuccess, nil
}

func mapDebtorsAsResponse(debtors []models.Debtor) *response.ListDebtorResponse {
	debtorsResponse := make([]response.DebtorResponse, len(debtors))

	for i, debtor := range debtors {
		debtorResponse := response.DebtorResponse{
			ID:                 debtor.ID.String(),
			UserID:             debtor.UserID.String(),
			NIK:                debtor.NIK,
			FullName:           debtor.FullName,
			LegalName:          debtor.LegalName,
			PlaceOfBirth:       debtor.PlaceOfBirth,
			DateOfBirth:        utils.FormatDate(debtor.DateOfBirth),
			Salary:             debtor.Salary,
			IdentityPictureUrl: debtor.IdentityPictureUrl,
			SelfiePictureUrl:   debtor.SelfiePictureUrl,
		}

		debtorsResponse[i] = debtorResponse
	}

	return &response.ListDebtorResponse{
		Data:  debtorsResponse,
		Total: len(debtors),
	}
}
