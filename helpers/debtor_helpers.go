package helpers

import (
	"credit/models"
	"credit/models/enums"
	"credit/utils"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	errOtrInvalidValue     = "otr value must be decimal"
	errTenorLimitNotEnough = "current limit is not enough"
)

type LoanCalculationResult struct {
	AdminFee      *big.Float
	TotalInterest *big.Float
	TotalLoan     *big.Float
}

func GenerateTransactionSequence(db *gorm.DB) (*string, *models.Sequence, error) {
	var sequence models.Sequence
	if err := db.Where("id = ?", "debtor_transaction").First(&sequence).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, nil, err
	}

	if sequence.ID == "" {
		sequence.ID = "debtor_transaction"
		sequence.Prefix = "TRANSACTION-"
		sequence.LastNumber = 1
	} else {
		sequence.LastNumber++
	}

	sequenceStr := fmt.Sprintf("%s%06d", sequence.Prefix, sequence.LastNumber)

	return &sequenceStr, &sequence, nil
}

// Simple tenor limit calculation by Salary
func GenerateDebtorTenorLimitBySalary(userID uuid.UUID, debtorID uuid.UUID, salary *big.Float) []models.DebtorTenorLimit {
	var debtorTenorLimits []models.DebtorTenorLimit

	tenor := []int{1, 2, 3, 6}

	loanToIncomeRatio := big.NewFloat(0.4)
	interestRate := big.NewFloat(0.12)

	// Loan Limit = (Salary * LoanToIncomeRatio) * ((1 + (InterestRate / 12)) ^ Month)
	for _, t := range tenor {
		loanIncome := new(big.Float).Mul(salary, loanToIncomeRatio)

		month := new(big.Float).SetInt(big.NewInt(int64(t)))

		monthlyInterestRate := new(big.Float).Add(big.NewFloat(1), new(big.Float).Quo(interestRate, big.NewFloat(12)))

		loanLimit := new(big.Float).Mul(loanIncome, monthlyInterestRate)
		loanLimit = new(big.Float).Mul(loanLimit, month)

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

// Simple loan calculation based on OTR
func GetLoanCalculationFromOTR(otr string) (*LoanCalculationResult, error) {
	otrDecimal, ok := new(big.Float).SetString(otr)
	if !ok {
		return nil, errors.New(errOtrInvalidValue)
	}

	// Calculate admin fee: OTR * 0.005
	adminFee := new(big.Float).Mul(otrDecimal, big.NewFloat(0.005))

	// Calculate total interest: 12% of OTR
	interestRate := big.NewFloat(0.12)
	totalInterest := new(big.Float).Mul(otrDecimal, interestRate)

	// Calculate total loan: OTR + adminFee
	totalLoan := new(big.Float)
	totalLoan.Add(otrDecimal, adminFee)
	totalLoan.Add(totalLoan, totalInterest)

	return &LoanCalculationResult{
		AdminFee:      adminFee,
		TotalInterest: totalInterest,
		TotalLoan:     totalLoan,
	}, nil
}

func ReduceDebtorTenor(db *gorm.DB, debtorTenorLimit *models.DebtorTenorLimit, loanAmount *big.Float) error {
	currentLimit, ok := new(big.Float).SetString(debtorTenorLimit.CurrentLimit)
	if !ok || currentLimit.Cmp(loanAmount) == -1 || currentLimit.Cmp(big.NewFloat(0)) == 0 {
		return errors.New(errTenorLimitNotEnough)
	}

	newCurrentLimit := new(big.Float).Sub(currentLimit, loanAmount)
	debtorTenorLimit.CurrentLimit = utils.FormatMoney(newCurrentLimit)

	return nil
}

func GenerateDebtorInstallment(db *gorm.DB, debtorTransaction models.DebtorTransaction, debtorTenorLimit models.DebtorTenorLimit, loanCalculationResult LoanCalculationResult) (*models.DebtorInstallment, error) {
	var totalMonth float64
	if debtorTenorLimit.TenorLimitType == enums.Monthly {
		totalMonth = float64(debtorTenorLimit.TenorDuration)
	} else {
		totalMonth = float64(debtorTenorLimit.TenorDuration) * float64(12)
	}

	monthlyInstallment := new(big.Float).Set(loanCalculationResult.TotalLoan).Quo(loanCalculationResult.TotalLoan, big.NewFloat(totalMonth))

	debtorInstallment := &models.DebtorInstallment{
		TenorLimitID:        debtorTenorLimit.ID,
		DebtorTransactionID: debtorTransaction.ID,
		MonthlyInstallment:  utils.FormatMoney(monthlyInstallment),
		TotalInstallment:    utils.FormatMoney(loanCalculationResult.TotalLoan),
		StartDatePeriod:     debtorTransaction.CreatedAt,
		EndDatePeriod:       debtorTransaction.CreatedAt.AddDate(0, debtorTenorLimit.TenorDuration, 0),
	}

	debtorInstallment.SetUser(debtorTransaction.UserID)
	debtorInstallment.SetAuditCreated(time.Now(), debtorTransaction.UserID)

	debtorInstallmentLines := make([]models.DebtorInstallmentLine, debtorTenorLimit.TenorDuration)
	for i := range debtorInstallmentLines {
		debtorInstallmentLine := models.DebtorInstallmentLine{
			DebtorInstallmentID: debtorInstallment.ID,
			InstallmentNumber:   i + 1,
			InstallmentAmount:   utils.FormatMoney(monthlyInstallment),
			DueDate:             debtorInstallment.StartDatePeriod.AddDate(0, i+1, 0),
			Status:              enums.Upcoming,
		}

		debtorInstallmentLine.SetUser(debtorTransaction.UserID)
		debtorInstallmentLine.SetAuditCreated(time.Now(), debtorTransaction.UserID)

		debtorInstallmentLines[i] = debtorInstallmentLine
	}

	debtorInstallment.InstallmentLines = debtorInstallmentLines
	return debtorInstallment, nil
}
