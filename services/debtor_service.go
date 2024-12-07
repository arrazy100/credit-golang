package services

import (
	"credit/dtos/request"
	"credit/dtos/response"
	custom_errors "credit/errors"
	"credit/models"
	"credit/models/enums"
	"credit/services/interfaces"
	"credit/utils"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var _ interfaces.DebtorInterface = (*DebtorService)(nil)

const (
	errDebtorAlreadyRegistered = "debtor already registered"
	errParseSalaryFailed       = "failed to parse salary %s"
	errParseDateFailed         = "failed to parse date %s"
	errDebtorNotRegistered     = "debtor not registered"
)

type DebtorService struct {
	db *gorm.DB
}

func NewDebtorService(db *gorm.DB) *DebtorService {
	return &DebtorService{db: db}
}

// TODO: Implement concurrency
func (s *DebtorService) RegisterDebtor(userID uuid.UUID, payload request.RegisterDebtorPayload) (*response.RegisterDebtorResponse, int, *custom_errors.ErrorValidation) {
	validationError := custom_errors.ValidateStruct(payload)
	if validationError != nil {
		return nil, statusBadRequest, validationError
	}

	var existingDebtor models.Debtor
	if err := s.db.Where("user_id = ?", userID).First(&existingDebtor).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, statusServerError, custom_errors.Convert(err)
	}

	if existingDebtor.ID != uuid.Nil {
		return nil, statusBadRequest, custom_errors.Convert(errors.New(errDebtorAlreadyRegistered))
	}

	salary, err := utils.ParseMoney(payload.Salary)
	if err != nil {
		return nil, statusBadRequest, custom_errors.Convert(errors.New(errParseSalaryFailed))
	}

	debtor, err := mapDebtorProfileWithPayload(payload)
	if err != nil {
		return nil, statusBadRequest, custom_errors.Convert(err)
	}

	debtor.SetUser(userID)
	debtor.SetAuditCreated(time.Now(), userID)
	debtor.TenorLimits = generateDebtorTenorLimitBySalary(userID, debtor.ID, salary)

	err = s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&debtor).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, statusServerError, custom_errors.Convert(err)
	}

	var savedDebtor models.Debtor
	if err := s.db.Preload("TenorLimits").Where("id = ?", debtor.ID).First(&savedDebtor).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, statusServerError, custom_errors.Convert(err)
	}

	return mapDebtorProfileAsResponse(*debtor), statusCreated, nil
}

func (s *DebtorService) DetailDebtor(userID uuid.UUID) (*response.DebtorResponse, int, *custom_errors.ErrorValidation) {
	var debtor models.Debtor
	if err := s.db.Where("user_id = ?", userID).First(&debtor).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, statusServerError, custom_errors.Convert(err)
	}

	if debtor.ID == uuid.Nil {
		return nil, statusBadRequest, custom_errors.Convert(errors.New(errDebtorNotRegistered))
	}

	return mapDebtorAsResponse(debtor), statusSuccess, nil
}

// Helpers

func mapDebtorProfileWithPayload(payload request.RegisterDebtorPayload) (*models.Debtor, error) {
	parsedDate, err := time.Parse(time.DateOnly, payload.DateOfBirth)
	if err != nil {
		return nil, fmt.Errorf(errParseDateFailed, payload.DateOfBirth)
	}

	return &models.Debtor{
		NIK:                payload.NIK,
		FullName:           payload.FullName,
		LegalName:          payload.LegalName,
		PlaceOfBirth:       payload.PlaceOfBirth,
		DateOfBirth:        parsedDate,
		Salary:             payload.Salary,
		IdentityPictureUrl: payload.IdentityPictureUrl,
		SelfiePictureUrl:   payload.SelfiePictureUrl,
	}, err
}

func mapDebtorProfileAsResponse(debtor models.Debtor) *response.RegisterDebtorResponse {
	debtorTenorLimits := make([]response.DebtorTenorLimitResponse, len(debtor.TenorLimits))
	for i, tenorLimit := range debtor.TenorLimits {
		tenorLimitResponse := response.DebtorTenorLimitResponse{
			TenorLimitType: tenorLimit.TenorLimitType,
			TenorDuration:  tenorLimit.TenorDuration,
			TotalLimit:     tenorLimit.TotalLimit,
			CurrentLimit:   tenorLimit.CurrentLimit,
		}

		debtorTenorLimits[i] = tenorLimitResponse
	}

	return &response.RegisterDebtorResponse{
		NIK:                debtor.NIK,
		FullName:           debtor.FullName,
		LegalName:          debtor.LegalName,
		PlaceOfBirth:       debtor.PlaceOfBirth,
		DateOfBirth:        utils.FormatDate(debtor.DateOfBirth),
		Salary:             debtor.Salary,
		IdentityPictureUrl: debtor.IdentityPictureUrl,
		SelfiePictureUrl:   debtor.SelfiePictureUrl,
		TenorLimits:        debtorTenorLimits,
	}
}

func mapDebtorAsResponse(debtor models.Debtor) *response.DebtorResponse {
	return &response.DebtorResponse{
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
}

// Simple tenor limit calculation by salary
func generateDebtorTenorLimitBySalary(userID uuid.UUID, debtorID uuid.UUID, salary *big.Float) []models.DebtorTenorLimit {
	var debtorTenorLimits []models.DebtorTenorLimit

	tenor := []int{1, 2, 3, 6}

	for _, t := range tenor {
		loanToIncomeRatio := big.NewFloat(0.4)
		interestRate := big.NewFloat(0.12)
		month := new(big.Float).SetInt(big.NewInt(int64(t)))

		// Calculate loan limit: (Salary * LoanToIncomeRatio) / (LoanTermMonths * (1 + interestRateDecimal))
		loanIncome := new(big.Float).Mul(salary, loanToIncomeRatio)
		onePlusInterest := new(big.Float).Add(big.NewFloat(1), interestRate)
		loanTermWithInterest := new(big.Float).Mul(month, onePlusInterest)
		loanLimit := new(big.Float).Quo(loanIncome, loanTermWithInterest)

		totalLimit := utils.FormatMoney(loanLimit)

		debtorTenorLimit := models.DebtorTenorLimit{
			DebtorID:       debtorID,
			TenorLimitType: enums.Monthly,
			TenorDuration:  t,
			TotalLimit:     totalLimit,
			CurrentLimit:   totalLimit,
		}
		debtorTenorLimit.SetUser(userID)
		debtorTenorLimit.SetAuditCreated(time.Now(), userID)

		debtorTenorLimits = append(debtorTenorLimits, debtorTenorLimit)
	}

	return debtorTenorLimits
}
