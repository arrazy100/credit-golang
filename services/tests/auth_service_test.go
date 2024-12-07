package services_tests

import (
	"credit/dtos/request"
	"credit/models/enums"
	"credit/services"
	"net/http"
)

func (suite *ServiceTestSuite) TestAuthService_Login() {
	service := services.NewAuthService(suite.Db)
	payload := request.LoginPayload{
		Email:    "admin@admin.com",
		Password: "admin",
	}
	response, status, err := service.Login(payload)

	suite.Nil(err)
	suite.Equal(status, http.StatusOK)
	suite.Equal(suite.Admin.ID, response.User.UserID)
	suite.Equal(suite.Admin.Role, response.User.Role)
}

func (suite *ServiceTestSuite) TestAuthService_LoginFailed() {
	service := services.NewAuthService(suite.Db)
	payload := request.LoginPayload{
		Email:    "admin2@admin.com",
		Password: "admin",
	}
	response, status, err := service.Login(payload)

	suite.Nil(response)
	suite.Equal(status, http.StatusBadRequest)
	suite.Equal(err.Message, "account with email admin2@admin.com not found")
}

func (suite *ServiceTestSuite) TestAuthService_LoginWrongPassword() {
	service := services.NewAuthService(suite.Db)
	payload := request.LoginPayload{
		Email:    "admin@admin.com",
		Password: "admin123",
	}
	_, status, err := service.Login(payload)
	suite.Equal(err.Message, "invalid password")
	suite.Equal(status, http.StatusUnauthorized)
}

func (suite *ServiceTestSuite) TestAuthService_Register() {
	service := services.NewAuthService(suite.Db)
	payload := request.RegisterPayload{
		Email:    "test@test.com",
		Password: "test123",
	}
	response, status, err := service.RegisterUser(payload)
	suite.Nil(err)
	suite.Equal(status, http.StatusCreated)
	suite.Equal(response.Email, payload.Email)
	suite.Equal(response.Role, enums.Debtor)
}

func (suite *ServiceTestSuite) TestAuthService_RegisterFailed() {
	service := services.NewAuthService(suite.Db)
	payload := request.RegisterPayload{
		Email:    "admin@admin.com",
		Password: "admin123",
	}
	response, status, err := service.RegisterUser(payload)
	suite.Nil(response)
	suite.Equal(status, http.StatusBadRequest)
	suite.Equal(err.Message, "email admin@admin.com already registered")
}
