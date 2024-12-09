package services

import (
	"credit/models"
	"credit/services"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestAdminService_ListDebtor(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	require.NoError(t, err)

	adminService := services.NewAdminService(gormDB)

	mockDebtors := []models.Debtor{
		{
			NIK:                "1234567890",
			FullName:           "John Doe",
			LegalName:          "John Doe",
			PlaceOfBirth:       "City X",
			DateOfBirth:        time.Now(),
			Salary:             "10000",
			IdentityPictureUrl: "http://example.com/identity.jpg",
			SelfiePictureUrl:   "http://example.com/selfie.jpg",
		},
		{
			NIK:                "0987654321",
			FullName:           "Jane Smith",
			LegalName:          "Jane Smith",
			PlaceOfBirth:       "City Y",
			DateOfBirth:        time.Now(),
			Salary:             "12000",
			IdentityPictureUrl: "http://example.com/identity2.jpg",
			SelfiePictureUrl:   "http://example.com/selfie2.jpg",
		},
	}

	for _, debtor := range mockDebtors {
		userID := uuid.New()
		debtor.SetUser(userID)
		debtor.SetAuditCreated(time.Now(), userID)
	}

	rows := sqlmock.NewRows([]string{"id", "user_id", "nik", "full_name", "legal_name", "place_of_birth", "date_of_birth", "salary", "identity_picture_url", "selfie_picture_url"}).
		AddRow(mockDebtors[0].ID, mockDebtors[0].UserID, mockDebtors[0].NIK, mockDebtors[0].FullName, mockDebtors[0].LegalName, mockDebtors[0].PlaceOfBirth, mockDebtors[0].DateOfBirth, mockDebtors[0].Salary, mockDebtors[0].IdentityPictureUrl, mockDebtors[0].SelfiePictureUrl).
		AddRow(mockDebtors[1].ID, mockDebtors[1].UserID, mockDebtors[1].NIK, mockDebtors[1].FullName, mockDebtors[1].LegalName, mockDebtors[1].PlaceOfBirth, mockDebtors[1].DateOfBirth, mockDebtors[1].Salary, mockDebtors[1].IdentityPictureUrl, mockDebtors[1].SelfiePictureUrl)

	mock.ExpectQuery(`SELECT (.+) FROM "debtors"`).
		WillReturnRows(rows)

	response, status, validationError := adminService.ListDebtor()

	assert.Nil(t, validationError)
	assert.Equal(t, 200, status)
	assert.Equal(t, 2, response.Total)
	assert.Len(t, response.Data, 2)

	assert.Equal(t, mockDebtors[0].ID.String(), response.Data[0].ID)
	assert.Equal(t, mockDebtors[0].FullName, response.Data[0].FullName)
	assert.Equal(t, mockDebtors[0].Salary, response.Data[0].Salary)

	assert.Equal(t, mockDebtors[1].ID.String(), response.Data[1].ID)
	assert.Equal(t, mockDebtors[1].FullName, response.Data[1].FullName)
	assert.Equal(t, mockDebtors[1].Salary, response.Data[1].Salary)
}
