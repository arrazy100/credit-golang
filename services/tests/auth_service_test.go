package services_tests

import (
	"credit/services"
)

func (suite *ServiceTestSuite) TestAuthService_Login() {
	authService := services.NewAuthService(suite.Db)
	response, err := authService.Login("admin@admin.com", "admin")

	suite.NoError(err)
	suite.Equal(suite.Admin.ID, response.User.UserID)
	suite.Equal(suite.Admin.Role, response.User.Role)
}

func (suite *ServiceTestSuite) TestAuthService_LoginFailed() {
	authService := services.NewAuthService(suite.Db)
	response, err := authService.Login("admin2@admin.com", "admin")
	suite.Error(err)
	suite.Nil(response)
}

func (suite *ServiceTestSuite) TestAuthService_LoginWrongPassword() {
	authService := services.NewAuthService(suite.Db)
	_, err := authService.Login("admin@admin.com", "admin123")
	suite.Error(err)
	suite.Equal(err.Error(), "invalid password")
}
