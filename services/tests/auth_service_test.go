package services_test

import (
	"credit/dtos/request"
	"credit/dtos/response"
	"credit/models/base"
	"credit/models/enums"
	"credit/services"
	"credit/utils"
	"net/http"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestAuthService_Login(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	require.NoError(t, err)

	authService := services.NewAuthService(gormDB)

	hashedPassword, _ := utils.HashPassword("password123")
	mockUser := base.User{
		ID:       uuid.New(),
		Email:    "test@example.com",
		Password: hashedPassword,
		Role:     enums.Admin,
	}

	t.Run("Success", func(t *testing.T) {
		mock.ExpectQuery(`SELECT (.+) FROM "users" WHERE (.+)`).
			WithArgs("test@example.com", 1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "email", "password", "role"}).
				AddRow(mockUser.ID, mockUser.Email, mockUser.Password, mockUser.Role))

		payload := request.LoginPayload{
			Email:    "test@example.com",
			Password: "password123",
		}

		resp, status, errValidation := authService.Login(payload)

		assert.Nil(t, errValidation)
		assert.Equal(t, 200, status)
		assert.Equal(t, mockUser.ID, resp.User.UserID)
		assert.Equal(t, mockUser.Role.String(), resp.User.Role)
	})

	t.Run("User Not Found", func(t *testing.T) {
		mock.ExpectQuery(`SELECT (.+) FROM "users" WHERE (.+)`).
			WithArgs("notfound@example.com", 1).
			WillReturnError(gorm.ErrRecordNotFound)

		payload := request.LoginPayload{
			Email:    "notfound@example.com",
			Password: "password123",
		}

		resp, status, errValidation := authService.Login(payload)

		assert.Nil(t, resp)
		assert.Equal(t, 400, status)
		assert.NotNil(t, errValidation)
	})

	t.Run("Invalid Password", func(t *testing.T) {
		mock.ExpectQuery(`SELECT (.+) FROM "users" WHERE (.+)`).
			WithArgs("test@example.com", 1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "email", "password", "role"}).
				AddRow(mockUser.ID, mockUser.Email, mockUser.Password, mockUser.Role))

		payload := request.LoginPayload{
			Email:    "test@example.com",
			Password: "wrongpassword",
		}

		resp, status, errValidation := authService.Login(payload)

		assert.Nil(t, resp)
		assert.Equal(t, 401, status)
		assert.NotNil(t, errValidation)
	})
}

func TestAuthService_RegisterUser(t *testing.T) {
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

	authService := services.NewAuthService(gormDB)

	payload := request.RegisterPayload{
		Email:    "test@example.com",
		Password: "password123",
	}

	mock.ExpectQuery(`SELECT (.+) FROM "users" WHERE (.+)`).
		WithArgs(payload.Email, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id"}))

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "users" (.+)`).
		WithArgs(sqlmock.AnyArg(), payload.Email, sqlmock.AnyArg(), enums.Debtor).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	resp, statusCode, errValidation := authService.RegisterUser(payload)

	if errValidation != nil {
		t.Errorf("Expected no validation error, but got: %v", errValidation)
	}

	if statusCode != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, statusCode)
	}

	expectedResponse := &response.RegisterResponse{
		UserID: resp.UserID,
		Email:  payload.Email,
		Role:   enums.Debtor.String(),
	}

	if resp.Email != expectedResponse.Email || resp.Role != expectedResponse.Role {
		t.Errorf("Expected response %v, but got %v", expectedResponse, resp)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
