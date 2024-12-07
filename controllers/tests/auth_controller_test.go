package controllers_tests

import (
	"bytes"
	"credit/controllers"
	"credit/dtos/request"
	"credit/dtos/response"
	custom_errors "credit/errors"
	"credit/models/enums"
	"credit/utils"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthController_LoginSuccess(t *testing.T) {
	userID := uuid.New()

	mockService := new(MockAuthService)
	mockService.On("Login", mock.Anything).Return(&response.LoginResponse{
		Token: "mocked_token",
		User: utils.SimpleAuth{
			UserID: userID,
			Role:   enums.Debtor,
		},
	}, http.StatusOK, nil)

	authController := controllers.NewAuthController(mockService)

	w, c := SetupRouter(authController)

	payload := request.LoginPayload{
		Email:    "user@user.com",
		Password: "test_password",
	}
	jsonPayload, _ := json.Marshal(payload)
	c.Request, _ = http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonPayload))

	authController.Login(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var loginResponse response.LoginResponse
	ParseBodyAsStruct(t, w.Body, &loginResponse)

	assert.Equal(t, "mocked_token", loginResponse.Token)
	assert.Equal(t, userID, loginResponse.User.UserID)
	assert.Equal(t, enums.Debtor, loginResponse.User.Role)

	mockService.AssertExpectations(t)
}

func TestAuthController_InvalidPayload(t *testing.T) {
	mockService := new(MockAuthService)
	authController := controllers.NewAuthController(mockService)

	w, c := SetupRouter(authController)

	c.Request, _ = http.NewRequest("POST", "/auth/login", bytes.NewBufferString("invalid_json"))

	authController.Login(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAuthController_InvalidPassword(t *testing.T) {
	field := "Password"

	mockService := new(MockAuthService)
	mockService.On("Login", mock.Anything).Return(nil, http.StatusUnauthorized,
		&custom_errors.ErrorValidation{
			Fields: []custom_errors.ErrorField{
				{Field: &field, Value: "Invalid password"},
			},
		},
	)

	authController := controllers.NewAuthController(mockService)

	w, c := SetupRouter(authController)

	payload := request.LoginPayload{
		Email:    "user@user.com",
		Password: "wrong_password",
	}
	jsonPayload, _ := json.Marshal(payload)
	c.Request, _ = http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonPayload))

	authController.Login(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	mockService.AssertExpectations(t)
}

func TestAuthController_RegisterUserSuccess(t *testing.T) {
	userID := uuid.New()

	mockService := new(MockAuthService)
	mockService.On("RegisterUser", mock.Anything).Return(&response.RegisterResponse{
		UserID: userID,
		Email:  "user@user.com",
		Role:   enums.Debtor,
	}, http.StatusCreated, nil)

	authController := controllers.NewAuthController(mockService)

	w, c := SetupRouter(authController)

	payload := request.RegisterPayload{
		Email:    "user@user.com",
		Password: "test_password",
	}

	jsonPayload, _ := json.Marshal(payload)
	c.Request, _ = http.NewRequest("POST", "/auth/register", bytes.NewBuffer(jsonPayload))

	authController.RegisterUser(c)

	assert.Equal(t, http.StatusCreated, w.Code)

	var registerResponse response.RegisterResponse
	ParseBodyAsStruct(t, w.Body, &registerResponse)

	assert.Equal(t, userID, registerResponse.UserID)
	assert.Equal(t, "user@user.com", registerResponse.Email)
	assert.Equal(t, enums.Debtor, registerResponse.Role)

	mockService.AssertExpectations(t)
}

func TestAuthController_RegisterUserInvalidPayload(t *testing.T) {
	mockService := new(MockAuthService)
	authController := controllers.NewAuthController(mockService)

	w, c := SetupRouter(authController)

	c.Request, _ = http.NewRequest("POST", "/auth/register", bytes.NewBufferString("invalid_json"))

	authController.RegisterUser(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	mockService.AssertExpectations(t)
}

func TestAuthController_RegisterUserValidationError(t *testing.T) {
	field := "Email"

	mockService := new(MockAuthService)
	mockService.On("RegisterUser", mock.Anything).Return(nil, http.StatusBadRequest,
		&custom_errors.ErrorValidation{
			Fields: []custom_errors.ErrorField{
				{Field: &field, Value: "Invalid email"},
			},
		},
	)

	authController := controllers.NewAuthController(mockService)

	w, c := SetupRouter(authController)

	payload := request.RegisterPayload{
		Email:    "invalid_email",
		Password: "test_password",
	}

	jsonPayload, _ := json.Marshal(payload)
	c.Request, _ = http.NewRequest("POST", "/auth/register", bytes.NewBuffer(jsonPayload))

	authController.RegisterUser(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	mockService.AssertExpectations(t)
}
