package services

import (
	"credit/dtos/request"
	"credit/dtos/response"
	"credit/helpers"
	"credit/models"
	"credit/models/enums"
	"credit/services/interfaces"
	"credit/utils"
	validations "credit/validations"
	"errors"
	"fmt"
	"log"
	"sort"
	"sync"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var _ interfaces.DebtorInterface = (*DebtorService)(nil)

const (
	errDebtorAlreadyRegistered          = "debtor already registered"
	errParseSalaryFailed                = "failed to parse salary %s"
	errParseDateFailed                  = "failed to parse date %s"
	errDebtorNotRegistered              = "debtor not registered"
	errParseUUIDFailed                  = "failed to parse uuid %s"
	errTenorLimitNotFound               = "tenor limit not found with id %s"
	errDebtorInstallmentLineNotFound    = "debtor installment line not found with id %s"
	errDebtorInstallmentLineAlreadyPaid = "debtor installment line already paid with id %s"
)

type DebtorService struct {
	db *gorm.DB
}

func NewDebtorService(db *gorm.DB) *DebtorService {
	return &DebtorService{db: db}
}

func (s *DebtorService) Register(userID uuid.UUID, payload request.RegisterDebtorPayload) (*response.RegisterDebtorResponse, int, *validations.ErrorValidation) {
	validationError := validations.ValidateStruct(payload)
	if validationError != nil {
		return nil, statusBadRequest, validationError
	}

	var existingDebtor models.Debtor
	if err := s.db.Where("user_id = ?", userID).First(&existingDebtor).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, statusServerError, validations.Convert(err)
	}

	if existingDebtor.ID != uuid.Nil {
		return nil, statusBadRequest, validations.Convert(errors.New(errDebtorAlreadyRegistered))
	}

	salary, err := utils.ParseMoney(payload.Salary)
	if err != nil {
		return nil, statusBadRequest, validations.Convert(errors.New(errParseSalaryFailed))
	}

	debtor, err := mapDebtorProfileWithPayload(payload)
	if err != nil {
		return nil, statusBadRequest, validations.Convert(err)
	}

	debtor.SetUser(userID)
	debtor.SetAuditCreated(time.Now(), userID)
	debtor.TenorLimits = helpers.GenerateDebtorTenorLimitBySalary(userID, debtor.ID, salary)

	err = s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&debtor).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, statusServerError, validations.Convert(err)
	}

	var savedDebtor models.Debtor
	if err := s.db.Preload("TenorLimits").Where("id = ?", debtor.ID).First(&savedDebtor).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, statusServerError, validations.Convert(err)
	}

	return mapDebtorProfileAsResponse(*debtor), statusCreated, nil
}

func (s *DebtorService) Detail(userID uuid.UUID) (*response.DebtorResponse, int, *validations.ErrorValidation) {
	var debtor models.Debtor
	if err := s.db.Where("user_id = ?", userID).Preload("TenorLimits").First(&debtor).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, statusServerError, validations.Convert(err)
	}

	if debtor.ID == uuid.Nil {
		return nil, statusBadRequest, validations.Convert(errors.New(errDebtorNotRegistered))
	}

	return mapDebtorAsResponse(debtor), statusSuccess, nil
}

func (s *DebtorService) CreateTransaction(userID uuid.UUID, payload request.DebtorTransactionPayload) (*response.DebtorTransactionResponse, int, *validations.ErrorValidation) {
	validationError := validations.ValidateStruct(payload)
	if validationError != nil {
		return nil, statusBadRequest, validationError
	}

	var debtor models.Debtor
	if err := s.db.Where("user_id = ?", userID).Preload("TenorLimits").First(&debtor).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, statusServerError, validations.Convert(err)
	}

	if debtor.ID == uuid.Nil {
		return nil, statusBadRequest, validations.Convert(errors.New(errDebtorNotRegistered))
	}

	transactionMappingResult, err := mapDebtorTransactionWithPayload(s.db, debtor, payload)
	if err != nil {
		return nil, statusBadRequest, validations.Convert(err)
	}

	debtorTransaction := transactionMappingResult.DebtorTransaction
	debtorTenorLimit := transactionMappingResult.DebtorTenorLimit
	debtorInstallment := transactionMappingResult.DebtorInstallment
	sequence := transactionMappingResult.Sequence

	debtorTransaction.Status = enums.Success // default to Success

	err = s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(debtorTransaction).Error; err != nil {
			return err
		}

		if err := tx.Save(debtorTenorLimit).Error; err != nil {
			return err
		}

		if err := tx.Create(debtorInstallment).Error; err != nil {
			return err
		}

		if err := tx.Save(sequence).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, statusServerError, validations.Convert(err)
	}

	return mapDebtorTransactionAsResponse(*debtorTransaction), statusCreated, nil
}

func (s *DebtorService) ListInstallment(userID uuid.UUID) (*response.ListDebtorInstallmentResponse, int, *validations.ErrorValidation) {
	var debtor models.Debtor
	if err := s.db.Where("user_id = ?", userID).First(&debtor).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, statusServerError, validations.Convert(err)
	}

	if debtor.ID == uuid.Nil {
		return nil, statusBadRequest, validations.Convert(errors.New(errDebtorNotRegistered))
	}

	var debtorInstallments []models.DebtorInstallment
	if err := s.db.Where("user_id = ?", userID).Preload("DebtorTransaction").Preload("DebtorTenorLimit").Preload("InstallmentLines").Find(&debtorInstallments).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, statusServerError, validations.Convert(err)
	}

	debtorInstallmentResponse := mapDebtorInstallmentsAsResponse(debtorInstallments)

	return &response.ListDebtorInstallmentResponse{
		Data:  debtorInstallmentResponse,
		Total: len(debtorInstallmentResponse),
	}, statusSuccess, nil
}

func (s *DebtorService) PayInstallmentLine(userID uuid.UUID, payload request.DebtorPayInstallmentLinePayload) (*response.DebtorInstallmentLineResponse, int, *validations.ErrorValidation) {
	validationError := validations.ValidateStruct(payload)
	if validationError != nil {
		return nil, statusBadRequest, validationError
	}

	var debtorInstallmentLine models.DebtorInstallmentLine
	if err := s.db.Where("user_id = ? AND id = ?", userID, payload.InstallmentLineID).
		Preload("DebtorInstallment").
		Preload("DebtorInstallment.DebtorTenorLimit").
		First(&debtorInstallmentLine).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, statusServerError, validations.Convert(err)
	}

	if debtorInstallmentLine.ID == uuid.Nil {
		return nil, statusBadRequest, validations.Convert(fmt.Errorf(errDebtorInstallmentLineNotFound, payload.InstallmentLineID))
	}

	if debtorInstallmentLine.Status == enums.Paid {
		return nil, statusBadRequest, validations.Convert(fmt.Errorf(errDebtorInstallmentLineAlreadyPaid, payload.InstallmentLineID))
	}

	currentDate := time.Now()

	debtorInstallmentLine.Status = enums.Paid
	debtorInstallmentLine.PaymentDate = &currentDate
	debtorInstallmentLine.SetAuditUpdated(currentDate, userID)

	debtorTenorLimit := debtorInstallmentLine.DebtorInstallment.DebtorTenorLimit

	currentLimitDecimal, err := utils.ParseMoney(debtorTenorLimit.CurrentLimit)
	if err != nil {
		return nil, statusServerError, validations.Convert(err)
	}

	installmentAmountDecimal, err := utils.ParseMoney(debtorInstallmentLine.InstallmentAmount)
	if err != nil {
		return nil, statusServerError, validations.Convert(err)
	}

	currentLimitDecimal.Add(currentLimitDecimal, installmentAmountDecimal)
	debtorTenorLimit.CurrentLimit = utils.FormatMoney(currentLimitDecimal)
	debtorTenorLimit.SetAuditUpdated(currentDate, userID)

	err = s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&debtorInstallmentLine).Error; err != nil {
			return err
		}

		if err := tx.Save(&debtorTenorLimit).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, statusServerError, validations.Convert(err)
	}

	return mapDebtorInstallmentLineAsResponse(debtorInstallmentLine), statusSuccess, nil
}

func (s *DebtorService) BatchUpdateOverdueInstallmentLine(wg *sync.WaitGroup, errCh chan<- error) {
	defer wg.Done()

	currentDateTime := time.Now().UTC()

	err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.
			Model(&models.DebtorInstallmentLine{}).
			Where("due_date < ? and status = ?", currentDateTime, enums.Upcoming).
			Update("status", enums.Overdue).Error; err != nil && err != gorm.ErrRecordNotFound {
			return err
		}

		return nil
	})

	if err != nil {
		errCh <- err
		return
	}

	log.Println("Successfully updated installment lines to overdue")
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
			TenorLimitType: tenorLimit.TenorLimitType.String(),
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
	debtorTenorLimits := make([]response.DebtorTenorLimitResponse, len(debtor.TenorLimits))
	for i, tenorLimit := range debtor.TenorLimits {
		tenorLimitResponse := response.DebtorTenorLimitResponse{
			ID:             tenorLimit.ID.String(),
			TenorLimitType: tenorLimit.TenorLimitType.String(),
			TenorDuration:  tenorLimit.TenorDuration,
			TotalLimit:     tenorLimit.TotalLimit,
			CurrentLimit:   tenorLimit.CurrentLimit,
		}

		debtorTenorLimits[i] = tenorLimitResponse
	}

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
		TenorLimits:        debtorTenorLimits,
	}
}

func mapDebtorTransactionWithPayload(db *gorm.DB, debtor models.Debtor, payload request.DebtorTransactionPayload) (*request.DebtorTransactionMapResult, error) {
	loanCalculationResult, err := helpers.GetLoanCalculationFromOTR(payload.OTR)
	if err != nil {
		return nil, err
	}

	var debtorTenorLimit models.DebtorTenorLimit
	for _, tenorLimit := range debtor.TenorLimits {
		if tenorLimit.ID == payload.TenorLimitID {
			debtorTenorLimit = tenorLimit
			break
		}
	}

	if debtorTenorLimit.ID == uuid.Nil {
		return nil, fmt.Errorf(errTenorLimitNotFound, payload.TenorLimitID)
	}

	err = helpers.ReduceDebtorTenor(db, &debtorTenorLimit, loanCalculationResult.TotalLoan)
	if err != nil {
		return nil, err
	}

	sequenceStr, sequence, err := helpers.GenerateTransactionSequence(db)
	if err != nil {
		return nil, err
	}

	debtorTransaction := models.DebtorTransaction{
		ContractNumber: *sequenceStr,
		OTR:            payload.OTR,
		AdminFee:       utils.FormatMoney(loanCalculationResult.AdminFee),
		TotalLoan:      utils.FormatMoney(loanCalculationResult.TotalLoan),
		TotalInterest:  utils.FormatMoney(loanCalculationResult.TotalInterest),
		AssetName:      payload.AssetName,
	}

	debtorTransaction.SetUser(debtor.UserID)
	debtorTransaction.SetAuditCreated(time.Now(), debtor.UserID)

	debtorInstallment, err := helpers.GenerateDebtorInstallment(db, debtorTransaction, debtorTenorLimit, *loanCalculationResult)
	if err != nil {
		return nil, err
	}

	return &request.DebtorTransactionMapResult{
		DebtorTransaction: &debtorTransaction,
		DebtorTenorLimit:  &debtorTenorLimit,
		DebtorInstallment: debtorInstallment,
		Sequence:          sequence,
	}, nil
}

func mapDebtorTransactionAsResponse(debtorTransaction models.DebtorTransaction) *response.DebtorTransactionResponse {
	return &response.DebtorTransactionResponse{
		ID:             debtorTransaction.ID.String(),
		ContractNumber: debtorTransaction.ContractNumber,
		OTR:            debtorTransaction.OTR,
		AdminFee:       debtorTransaction.AdminFee,
		TotalLoan:      debtorTransaction.TotalLoan,
		TotalInterest:  debtorTransaction.TotalInterest,
		AssetName:      debtorTransaction.AssetName,
		Status:         debtorTransaction.Status.String(),
	}
}

func mapDebtorInstallmentsAsResponse(debtorInstallment []models.DebtorInstallment) []response.DebtorInstallmentResponse {
	debtorInstallmentResponse := make([]response.DebtorInstallmentResponse, len(debtorInstallment))
	for i, installment := range debtorInstallment {
		debtorInstallmentLineResponse := make([]response.DebtorInstallmentLineResponse, len(installment.InstallmentLines))
		installmentLinePaidCount := 0
		for j, installmentLine := range installment.InstallmentLines {
			if installmentLine.Status == enums.Paid {
				installmentLinePaidCount++
			}

			debtorInstallmentLineResponse[j] = *mapDebtorInstallmentLineAsResponse(installmentLine)
		}

		sort.Sort(response.ByInstallmentNumber(debtorInstallmentLineResponse))

		debtorInstallmentResponse[i] = response.DebtorInstallmentResponse{
			ID:                    installment.ID.String(),
			DebtorTransaction:     *mapDebtorTransactionAsResponse(installment.DebtorTransaction),
			TenorDuration:         installment.DebtorTenorLimit.TenorDuration,
			TotalInstallmentCount: len(installment.InstallmentLines),
			PaidInstallmentCount:  installmentLinePaidCount,
			MonthlyInstallment:    installment.MonthlyInstallment,
			StartDatePeriod:       utils.FormatDate(installment.StartDatePeriod),
			EndDatePeriod:         utils.FormatDate(installment.EndDatePeriod),
			Lines:                 debtorInstallmentLineResponse,
		}
	}

	return debtorInstallmentResponse
}

func mapDebtorInstallmentLineAsResponse(installmentLine models.DebtorInstallmentLine) *response.DebtorInstallmentLineResponse {
	var paymentDateStr *string
	paymentDate := installmentLine.PaymentDate
	if paymentDate != nil {
		formatted := utils.FormatDate(*installmentLine.PaymentDate)
		paymentDateStr = &formatted
	}

	return &response.DebtorInstallmentLineResponse{
		ID:                installmentLine.ID.String(),
		DueDate:           utils.FormatDate(installmentLine.DueDate),
		InstallmentNumber: installmentLine.InstallmentNumber,
		InstallmentAmount: installmentLine.InstallmentAmount,
		PaymentDate:       paymentDateStr,
		Status:            installmentLine.Status.String(),
	}
}
