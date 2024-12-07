package services_tests

import (
	"credit/dtos/request"
	"credit/services"
)

func (suite *ServiceTestSuite) TestDebtorService_RegisterDebtor() {
	service := services.NewDebtorService(suite.Db)

	payload := request.RegisterDebtorPayload{
		NIK:                "12345678901234567812345",
		FullName:           "John Doe",
		LegalName:          "John Doe Corporation",
		PlaceOfBirth:       "Jakarta",
		DateOfBirth:        "1990-01-01",
		Salary:             "5000000",
		IdentityPictureUrl: "https://example.com/identity.jpg",
		SelfiePictureUrl:   "https://example.com/selfie.jpg",
	}

	response, status, err := service.RegisterDebtor(suite.DebtorUser.ID, payload)

	suite.Nil(err)
	suite.Equal(201, status)
	suite.NotNil(response)
	suite.Equal(payload.NIK, response.NIK)
	suite.Equal(payload.FullName, response.FullName)
	suite.Equal(payload.LegalName, response.LegalName)
	suite.Equal(payload.PlaceOfBirth, response.PlaceOfBirth)
	suite.Equal(payload.DateOfBirth, response.DateOfBirth)
	suite.Equal(payload.Salary, response.Salary)
	suite.Equal(payload.IdentityPictureUrl, response.IdentityPictureUrl)
	suite.Equal(payload.SelfiePictureUrl, response.SelfiePictureUrl)
	suite.Equal(4, len(response.TenorLimits))

	response, status, err = service.RegisterDebtor(suite.DebtorUser.ID, payload)

	suite.NotNil(err)
	suite.Equal(400, status)
	suite.Nil(response)
}

func (suite *ServiceTestSuite) TestDebtorService_RegisterDebtorValidationError() {
	service := services.NewDebtorService(suite.Db)

	payload := request.RegisterDebtorPayload{}

	response, status, err := service.RegisterDebtor(suite.DebtorUser.ID, payload)

	suite.NotNil(err)
	suite.Equal(400, status)
	suite.Nil(response)
}

func (suite *ServiceTestSuite) TestDebtorService_RegisterDebtorInvalidSalary() {
	service := services.NewDebtorService(suite.Db)

	payload := request.RegisterDebtorPayload{
		NIK:                "12345678901234567812345",
		FullName:           "John Doe",
		LegalName:          "John Doe Corporation",
		PlaceOfBirth:       "Jakarta",
		DateOfBirth:        "1990-01-01",
		Salary:             "invalid_salary",
		IdentityPictureUrl: "https://example.com/identity.jpg",
		SelfiePictureUrl:   "https://example.com/selfie.jpg",
	}

	response, status, err := service.RegisterDebtor(suite.DebtorUser.ID, payload)

	suite.NotNil(err)
	suite.Equal(400, status)
	suite.Nil(response)
}
