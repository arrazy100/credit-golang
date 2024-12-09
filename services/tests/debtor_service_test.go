package services

import (
	"credit/dtos/request"
	"credit/dtos/response"
	"credit/models/enums"
	"credit/services"
	"credit/utils"
	"net/http"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestDebtorService_Register(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database connection", err)
	}

	expectedUserID := uuid.New()
	mock.ExpectQuery(`SELECT (.+) FROM "debtors" WHERE (.+)`).
		WithArgs(expectedUserID, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id"}))

	mock.ExpectBegin()

	date := time.Date(1990, 01, 01, 0, 0, 0, 0, time.UTC)
	mock.ExpectQuery(`INSERT INTO "debtors" (.+)`).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, expectedUserID, "123456789011121314", "John Doe", "John Doe Legal", "New York", date, "50000", "http://example.com/id.jpg", "http://example.com/selfie.jpg", sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"created_at", "updated_at"}).AddRow(time.Now(), time.Now()))

	mock.ExpectQuery(`INSERT INTO "debtor_tenor_limits" (.+)`).
		WillReturnRows(sqlmock.NewRows([]string{"created_at", "updated_at"}).AddRow(time.Now(), time.Now()))

	mock.ExpectCommit()

	rows := sqlmock.NewRows([]string{"id", "user_id", "nik", "full_name", "legal_name", "place_of_birth", "date_of_birth", "salary", "identity_picture_url", "selfie_picture_url"}).
		AddRow(uuid.New(), expectedUserID, "123456789011121314", "John Doe", "John Doe", "New York", date, "50000", "http://example.com/id.jpg", "http://example.com/selfie.jpg")
	mock.ExpectQuery(`SELECT (.+) FROM "debtors" WHERE (.+)`).
		WillReturnRows(rows)

	tenorLimits := sqlmock.NewRows([]string{"id", "tenor_limit_type", "tenor_duration", "total_limit", "current_limit", "debtor_id"})
	mock.ExpectQuery(`SELECT (.+) FROM "debtor_tenor_limits" WHERE (.+)`).
		WillReturnRows(tenorLimits)

	service := services.NewDebtorService(gormDB)

	payload := request.RegisterDebtorPayload{
		NIK:                "123456789011121314",
		FullName:           "John Doe",
		LegalName:          "John Doe Legal",
		PlaceOfBirth:       "New York",
		DateOfBirth:        "1990-01-01",
		Salary:             "50000",
		IdentityPictureUrl: "http://example.com/id.jpg",
		SelfiePictureUrl:   "http://example.com/selfie.jpg",
	}

	resp, statusCode, errValidation := service.Register(expectedUserID, payload)

	if errValidation != nil {
		t.Errorf("Expected no error, but got: %v", *errValidation)
	}

	if statusCode != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, statusCode)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	expectedResponse := &response.RegisterDebtorResponse{
		NIK:                "123456789011121314",
		FullName:           "John Doe",
		LegalName:          "John Doe Legal",
		PlaceOfBirth:       "New York",
		DateOfBirth:        "1990-01-01",
		Salary:             "50000",
		IdentityPictureUrl: "http://example.com/id.jpg",
		SelfiePictureUrl:   "http://example.com/selfie.jpg",
	}

	if !compareRegisterDebtorResponses(expectedResponse, resp) {
		t.Errorf("Expected response %v, but got %v", expectedResponse, resp)
	}
}

func TestDebtorService_Detail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database connection", err)
	}

	dueDate := time.Now()
	expectedID := uuid.New()
	expectedUserID := uuid.New()
	rows := sqlmock.NewRows([]string{"id", "user_id", "nik", "full_name", "legal_name", "place_of_birth", "date_of_birth", "salary", "identity_picture_url", "selfie_picture_url"}).
		AddRow(expectedID, expectedUserID, "1234567890", "John Doe", "John Doe", "New York", dueDate, "50000", "http://example.com/id.jpg", "http://example.com/selfie.jpg")
	mock.ExpectQuery(`SELECT (.+) FROM "debtors" WHERE (.+)`).
		WithArgs(expectedUserID, 1).
		WillReturnRows(rows)

	tenorLimitId := uuid.New()
	tenorLimits := sqlmock.NewRows([]string{"id", "tenor_limit_type", "tenor_duration", "total_limit", "current_limit", "debtor_id"}).
		AddRow(tenorLimitId, enums.Monthly, 1, 1, 1, expectedID)
	mock.ExpectQuery(`SELECT (.+) FROM "debtor_tenor_limits" WHERE (.+)`).
		WithArgs(expectedID).
		WillReturnRows(tenorLimits)

	service := services.NewDebtorService(gormDB)

	resp, statusCode, errValidation := service.Detail(expectedUserID)

	if errValidation != nil {
		t.Errorf("Expected no error, but got: %v", errValidation)
	}

	if statusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, statusCode)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	expectedResponse := &response.DebtorResponse{
		ID:                 resp.ID,
		UserID:             resp.UserID,
		NIK:                "1234567890",
		FullName:           "John Doe",
		LegalName:          "John Doe",
		PlaceOfBirth:       "New York",
		DateOfBirth:        utils.FormatDate(dueDate),
		Salary:             "50000",
		IdentityPictureUrl: "http://example.com/id.jpg",
		SelfiePictureUrl:   "http://example.com/selfie.jpg",
		TenorLimits: []response.DebtorTenorLimitResponse{
			{
				ID:             tenorLimitId.String(),
				TenorLimitType: enums.Monthly.String(),
				TenorDuration:  1,
				TotalLimit:     "1",
				CurrentLimit:   "1",
			},
		},
	}

	if !compareDebtorResponses(expectedResponse, resp) {
		t.Errorf("Expected response %v, but got %v", expectedResponse, resp)
	}
}

func TestDebtorService_CreateTransaction(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database connection", err)
	}

	expectedID := uuid.New()
	expectedUserID := uuid.New()
	rows := sqlmock.NewRows([]string{"id", "user_id", "nik", "full_name", "legal_name", "place_of_birth", "date_of_birth", "salary", "identity_picture_url", "selfie_picture_url"}).
		AddRow(expectedID, expectedUserID, "1234567890", "John Doe", "John Doe Legal", "New York", time.Now(), "50000", "http://example.com/id.jpg", "http://example.com/selfie.jpg")
	mock.ExpectQuery(`SELECT (.+) FROM "debtors" WHERE (.+)`).
		WithArgs(expectedUserID, 1).
		WillReturnRows(rows)

	tenorLimitId := uuid.New()
	tenorLimits := sqlmock.NewRows([]string{"id", "tenor_limit_type", "tenor_duration", "total_limit", "current_limit", "debtor_id"}).
		AddRow(tenorLimitId, enums.Monthly, 1, 20000, 20000, expectedID)
	mock.ExpectQuery(`SELECT (.+) FROM "debtor_tenor_limits" WHERE (.+)`).
		WithArgs(expectedID).
		WillReturnRows(tenorLimits)

	mock.ExpectQuery(`SELECT (.+) FROM "sequences" WHERE (.+)`).
		WithArgs("debtor_transaction", 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "prefix", "last_number"}))

	mock.ExpectBegin()

	mock.ExpectQuery(`INSERT INTO "debtor_transactions" (.+) RETURNING "created_at","updated_at"`).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, expectedUserID, "TRANSACTION-000001", "10000", "50.00", "11250.00", "1200.00", "Car", enums.Success, sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"created_at", "updated_at"}).AddRow(time.Now(), time.Now()))

	mock.ExpectExec(`UPDATE "debtor_tenor_limits" SET (.+) WHERE (.+)`).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectQuery(`INSERT INTO "debtor_installments" (.+) RETURNING "created_at","updated_at"`).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, expectedUserID, sqlmock.AnyArg(), tenorLimitId, "11250.00", "11250.00", sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"created_at", "updated_at"}).AddRow(time.Now(), time.Now()))

	mock.ExpectQuery(`INSERT INTO "debtor_installment_lines" (.+) ON CONFLICT (.+) DO UPDATE SET (.+) RETURNING "created_at","updated_at"`).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, expectedUserID, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), nil, enums.Upcoming, sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"created_at", "updated_at"}).AddRow(time.Now(), time.Now()))

	mock.ExpectExec(`UPDATE "sequences" SET (.+) WHERE (.+)`).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	service := services.NewDebtorService(gormDB)

	payload := request.DebtorTransactionPayload{
		OTR:          "10000",
		AssetName:    "Car",
		TenorLimitID: tenorLimitId,
	}
	resp, statusCode, errValidation := service.CreateTransaction(expectedUserID, payload)

	if errValidation != nil {
		t.Errorf("Expected no error, but got: %v", errValidation)
	}

	if statusCode != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, statusCode)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	expectedResponse := &response.DebtorTransactionResponse{
		ID:             resp.ID,
		ContractNumber: resp.ContractNumber,
		OTR:            "10000",
		AdminFee:       resp.AdminFee,
		TotalLoan:      resp.TotalLoan,
		TotalInterest:  resp.TotalInterest,
		AssetName:      "Car",
		Status:         "Success",
	}
	if !compareDebtorTransactionResponses(expectedResponse, resp) {
		t.Errorf("Expected response %v, but got %v", expectedResponse, resp)
	}
}

func TestDebtorService_ListInstallment(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database connection", err)
	}

	expectedUserID := uuid.New()
	debtorRows := sqlmock.NewRows([]string{"id", "user_id", "nik", "full_name", "legal_name", "place_of_birth", "date_of_birth", "salary", "identity_picture_url", "selfie_picture_url"}).
		AddRow(uuid.New(), expectedUserID, "1234567890", "John Doe", "John Doe Legal", "New York", time.Now(), "50000", "http://example.com/id.jpg", "http://example.com/selfie.jpg")

	mock.ExpectQuery(`SELECT (.+) FROM "debtors" WHERE (.+)$`).
		WillReturnRows(debtorRows)

	debtorInstallmentID := uuid.New()
	debtorInstallmentRows := sqlmock.NewRows([]string{"id", "user_id", "current_limit"}).
		AddRow(debtorInstallmentID, expectedUserID, "20000")
	mock.ExpectQuery(`SELECT (.+) FROM "debtor_installments" WHERE (.+)$`).
		WillReturnRows(debtorInstallmentRows)

	debtorInstallmentLineRows := sqlmock.NewRows([]string{"id", "user_id", "debtor_installment_id", "due_date", "installment_number", "installment_amount", "payment_date", "status"}).
		AddRow(uuid.New(), expectedUserID, debtorInstallmentID, time.Now(), 1, 10000, nil, enums.Upcoming)
	mock.ExpectQuery(`SELECT (.+) FROM "debtor_installment_lines" WHERE (.+)$`).
		WillReturnRows(debtorInstallmentLineRows)

	service := services.NewDebtorService(gormDB)

	resp, statusCode, errValidation := service.ListInstallment(expectedUserID)

	if errValidation != nil {
		t.Errorf("Expected no error, but got: %v", errValidation)
	}

	if statusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, statusCode)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	expectedResponse := &response.ListDebtorInstallmentResponse{
		Data:  resp.Data,
		Total: 1,
	}

	if !compareListDebtorInstallmentResponses(expectedResponse, resp) {
		t.Errorf("Expected response %v, but got %v", expectedResponse, resp)
	}
}

func TestDebtorService_PayInstallmentLine(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database connection", err)
	}

	dueDate := time.Now()
	expectedUserID := uuid.New()
	installmentLineID := uuid.New()
	installmentID := uuid.New()
	rows := sqlmock.NewRows([]string{"id", "user_id", "debtor_installment_id", "installment_number", "installment_amount", "due_date", "status"}).
		AddRow(installmentLineID, expectedUserID, installmentID, 1, "1000", dueDate, enums.Upcoming)
	mock.ExpectQuery(`SELECT (.+) FROM "debtor_installment_lines" WHERE (.+)`).
		WithArgs(expectedUserID, installmentLineID, 1).
		WillReturnRows(rows)

	tenorLimitID := uuid.New()
	debtorTransactionID := uuid.New()
	startDatePeriod := time.Now()
	endDatePeriod := time.Now()
	mock.ExpectQuery(`SELECT (.+) FROM "debtor_installments" WHERE (.+)`).
		WithArgs(installmentID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_by_id", "updated_by_id", "user_id", "tenor_limit_id", "debtor_transaction_id", "monthly_installment", "total_installment", "start_date_period", "end_date_period"}).
			AddRow(installmentID, expectedUserID, nil, expectedUserID, tenorLimitID, debtorTransactionID, "1000", "1000", startDatePeriod, endDatePeriod))

	debtorID := uuid.New()
	mock.ExpectQuery(`SELECT (.+) FROM "debtor_tenor_limits" WHERE (.+)`).
		WithArgs(tenorLimitID).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "created_by_id", "updated_by_id", "user_id", "debtor_id", "tenor_limit_type", "tenor_duration", "total_limit", "current_limit",
		}).AddRow(tenorLimitID, expectedUserID, nil, expectedUserID, debtorID, enums.Monthly, 1, "5000000", "5000000"))

	mock.ExpectBegin()

	mock.ExpectQuery(`INSERT INTO "debtor_tenor_limits" (.+) VALUES (.+) ON CONFLICT DO NOTHING RETURNING (.+)`).
		WithArgs(tenorLimitID, expectedUserID, nil, expectedUserID, debtorID, enums.Monthly, 1, "5000000", "5000000").
		WillReturnRows(sqlmock.NewRows([]string{"created_at", "updated_at"}).AddRow(time.Now(), time.Now()))

	mock.ExpectQuery(`INSERT INTO "debtor_installments" (.+) VALUES (.+) ON CONFLICT DO NOTHING RETURNING (.+)`).
		WithArgs(installmentID, expectedUserID, nil, expectedUserID, debtorTransactionID, tenorLimitID, "1000", "1000", startDatePeriod, endDatePeriod).
		WillReturnRows(sqlmock.NewRows([]string{"created_at", "updated_at"}).AddRow(time.Now(), time.Now()))

	mock.ExpectExec(`UPDATE "debtor_installment_lines"`).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec(`UPDATE "debtor_tenor_limits"`).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	service := services.NewDebtorService(gormDB)

	payload := request.DebtorPayInstallmentLinePayload{
		InstallmentLineID: installmentLineID,
	}

	resp, statusCode, errValidation := service.PayInstallmentLine(expectedUserID, payload)

	if errValidation != nil {
		t.Errorf("Expected no error, but got: %v", errValidation)
	}

	if statusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, statusCode)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	expectedResponse := &response.DebtorInstallmentLineResponse{
		ID:                resp.ID,
		DueDate:           utils.FormatDate(dueDate),
		InstallmentNumber: 1,
		InstallmentAmount: "1000",
		PaymentDate:       resp.PaymentDate,
		Status:            "Paid",
	}

	if !compareDebtorInstallmentLineResponses(expectedResponse, resp) {
		t.Errorf("Expected response %v, but got %v", expectedResponse, resp)
	}
}

// Helper functions for comparison

func compareRegisterDebtorResponses(expected, actual *response.RegisterDebtorResponse) bool {
	if expected.NIK != actual.NIK ||
		expected.FullName != actual.FullName ||
		expected.LegalName != actual.LegalName ||
		expected.PlaceOfBirth != actual.PlaceOfBirth ||
		expected.DateOfBirth != actual.DateOfBirth ||
		expected.Salary != actual.Salary ||
		expected.IdentityPictureUrl != actual.IdentityPictureUrl ||
		expected.SelfiePictureUrl != actual.SelfiePictureUrl {
		return false
	}

	return true
}

func compareDebtorResponses(expected, actual *response.DebtorResponse) bool {
	if expected.ID != actual.ID ||
		expected.UserID != actual.UserID ||
		expected.NIK != actual.NIK ||
		expected.FullName != actual.FullName ||
		expected.LegalName != actual.LegalName ||
		expected.PlaceOfBirth != actual.PlaceOfBirth ||
		expected.DateOfBirth != actual.DateOfBirth ||
		expected.Salary != actual.Salary ||
		expected.IdentityPictureUrl != actual.IdentityPictureUrl ||
		expected.SelfiePictureUrl != actual.SelfiePictureUrl {
		return false
	}

	if len(expected.TenorLimits) != len(actual.TenorLimits) {
		return false
	}

	for i := range expected.TenorLimits {
		if expected.TenorLimits[i] != actual.TenorLimits[i] {
			return false
		}
	}

	return true
}

func compareDebtorTransactionResponses(expected, actual *response.DebtorTransactionResponse) bool {
	if expected.ID != actual.ID ||
		expected.ContractNumber != actual.ContractNumber ||
		expected.OTR != actual.OTR ||
		expected.AdminFee != actual.AdminFee ||
		expected.TotalLoan != actual.TotalLoan ||
		expected.TotalInterest != actual.TotalInterest ||
		expected.AssetName != actual.AssetName ||
		expected.Status != actual.Status {
		return false
	}

	return true
}

func compareListDebtorInstallmentResponses(expected, actual *response.ListDebtorInstallmentResponse) bool {
	if expected.Total != actual.Total {
		return false
	}

	if len(expected.Data) != len(actual.Data) {
		return false
	}

	for i := range expected.Data {
		if !compareDebtorInstallmentResponses(&expected.Data[i], &actual.Data[i]) {
			return false
		}
	}

	return true
}

func compareDebtorInstallmentResponses(expected, actual *response.DebtorInstallmentResponse) bool {
	if expected.ID != actual.ID ||
		expected.DebtorTransaction != actual.DebtorTransaction ||
		expected.TenorDuration != actual.TenorDuration ||
		expected.TotalInstallmentCount != actual.TotalInstallmentCount ||
		expected.PaidInstallmentCount != actual.PaidInstallmentCount ||
		expected.MonthlyInstallment != actual.MonthlyInstallment ||
		expected.StartDatePeriod != actual.StartDatePeriod ||
		expected.EndDatePeriod != actual.EndDatePeriod {
		return false
	}

	if len(expected.Lines) != len(actual.Lines) {
		return false
	}

	for i := range expected.Lines {
		if !compareDebtorInstallmentLineResponses(&expected.Lines[i], &actual.Lines[i]) {
			return false
		}
	}

	return true
}

func compareDebtorInstallmentLineResponses(expected, actual *response.DebtorInstallmentLineResponse) bool {
	if expected.ID != actual.ID ||
		expected.DueDate != actual.DueDate ||
		expected.InstallmentNumber != actual.InstallmentNumber ||
		expected.InstallmentAmount != actual.InstallmentAmount ||
		expected.PaymentDate != actual.PaymentDate ||
		expected.Status != actual.Status {
		return false
	}

	return true
}
