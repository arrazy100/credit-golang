package controllers_tests

import (
	"bytes"
	"credit/controllers"
	"credit/dtos/request"
	"credit/dtos/response"
	"credit/models"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (suite *ControllerTestSuite) TestDebtorController_RegisterDebtor() {
	controller := controllers.NewDebtorController(suite.Service.DebtorService)
	endpoint := "/debtor/register"
	method := "POST"

	// Success
	{
		payload := request.RegisterDebtorPayload{
			NIK:                "3232323232323232",
			FullName:           "Andi Pratama",
			LegalName:          "Andi Pratama",
			PlaceOfBirth:       "Jakarta",
			DateOfBirth:        "1990-03-03",
			Salary:             "6000000",
			IdentityPictureUrl: "http://identity-andi.com",
			SelfiePictureUrl:   "http://selfie-andi.com",
		}
		jsonPayload, _ := json.Marshal(payload)

		w, c := SetupRouter(controller)
		c.Request, _ = http.NewRequest(method, endpoint, bytes.NewBuffer(jsonPayload))
		c.Request.Header.Set("Authorization", suite.DebtorToken)

		controller.Register(c)

		var registerDebtorResponse response.RegisterDebtorResponse
		ParseBodyAsStruct(nil, w.Body, &registerDebtorResponse)

		suite.Equal(http.StatusCreated, w.Code)
		suite.Equal(registerDebtorResponse.NIK, payload.NIK)
		suite.Equal(4, len(registerDebtorResponse.TenorLimits))
	}

	// Empty Token
	{
		w, c := SetupRouter(controller)
		c.Request, _ = http.NewRequest(method, endpoint, &bytes.Buffer{})

		controller.Register(c)

		suite.Equal(http.StatusUnauthorized, w.Code)
	}

	// Invalid Payload
	{
		w, c := SetupRouter(controller)
		c.Request, _ = http.NewRequest(method, endpoint, &bytes.Buffer{})
		c.Request.Header.Set("Authorization", suite.DebtorToken)

		controller.Register(c)

		suite.Equal(http.StatusBadRequest, w.Code)
	}

	// Unauthorized
	{
		w, c := SetupRouter(controller)
		c.Request, _ = http.NewRequest(method, endpoint, nil)
		c.Request.Header.Set("Authorization", suite.AdminToken)

		controller.Detail(c)

		suite.Equal(http.StatusUnauthorized, w.Code)
	}

	// Already Registered
	{
		payload := request.RegisterDebtorPayload{
			NIK:                "3232323232323232",
			FullName:           "Andi Pratama",
			LegalName:          "Andi Pratama",
			PlaceOfBirth:       "Jakarta",
			DateOfBirth:        "1990-03-03",
			Salary:             "6000000",
			IdentityPictureUrl: "http://identity-andi.com",
			SelfiePictureUrl:   "http://selfie-andi.com",
		}
		jsonPayload, _ := json.Marshal(payload)

		w, c := SetupRouter(controller)
		c.Request, _ = http.NewRequest(method, endpoint, bytes.NewBuffer(jsonPayload))
		c.Request.Header.Set("Authorization", suite.TestUserToken)

		controller.Register(c)

		suite.Equal(http.StatusBadRequest, w.Code)
	}
}

func (suite *ControllerTestSuite) TestDebtorController_DetailDebtor() {
	controller := controllers.NewDebtorController(suite.Service.DebtorService)
	endpoint := "/debtor/detail"
	method := "GET"

	// Success
	{
		w, c := SetupRouter(controller)
		c.Request, _ = http.NewRequest(method, endpoint, nil)
		c.Request.Header.Set("Authorization", suite.TestUserToken)

		controller.Detail(c)

		var detailDebtorResponse response.DebtorResponse
		ParseBodyAsStruct(nil, w.Body, &detailDebtorResponse)

		suite.Equal(http.StatusOK, w.Code)
		suite.Equal(detailDebtorResponse.NIK, suite.RegisteredDebtor.NIK)
		suite.Equal(4, len(detailDebtorResponse.TenorLimits))
	}

	// Empty Token
	{
		w, c := SetupRouter(controller)
		c.Request, _ = http.NewRequest(method, endpoint, nil)

		controller.Detail(c)

		suite.Equal(http.StatusUnauthorized, w.Code)
	}

	// Unauthorized
	{
		w, c := SetupRouter(controller)
		c.Request, _ = http.NewRequest(method, endpoint, nil)
		c.Request.Header.Set("Authorization", suite.AdminToken)

		controller.Detail(c)

		suite.Equal(http.StatusUnauthorized, w.Code)
	}

	// Not Registered
	{
		w, c := SetupRouter(controller)
		c.Request, _ = http.NewRequest(method, endpoint, nil)
		c.Request.Header.Set("Authorization", suite.DebtorToken)

		controller.Detail(c)

		suite.Equal(http.StatusBadRequest, w.Code)
	}
}

func (suite *ControllerTestSuite) TestDebtorController_CreateTransaction() {
	controller := controllers.NewDebtorController(suite.Service.DebtorService)
	endpoint := "/debtor/transaction"
	method := "POST"

	var tenorLimit models.DebtorTenorLimit
	suite.Db.First(&tenorLimit).Where("user_id = ?", suite.TestUser.ID)

	// Success
	{
		payload := request.DebtorTransactionPayload{
			OTR:          "1000000",
			AssetName:    "Asset",
			TenorLimitID: tenorLimit.ID,
		}
		jsonPayload, _ := json.Marshal(payload)

		w, c := SetupRouter(controller)
		c.Request, _ = http.NewRequest(method, endpoint, bytes.NewBuffer(jsonPayload))
		c.Request.Header.Set("Authorization", suite.TestUserToken)

		controller.CreateTransaction(c)

		var transactionResponse response.DebtorTransactionResponse
		ParseBodyAsStruct(nil, w.Body, &transactionResponse)

		suite.Equal(http.StatusCreated, w.Code)
		suite.Equal(payload.AssetName, transactionResponse.AssetName)
		suite.Equal("Success", transactionResponse.Status)
	}

	// Empty Token
	{
		w, c := SetupRouter(controller)
		c.Request, _ = http.NewRequest(method, endpoint, nil)

		controller.CreateTransaction(c)

		suite.Equal(http.StatusUnauthorized, w.Code)
	}

	// Unauthorized
	{
		w, c := SetupRouter(controller)
		c.Request, _ = http.NewRequest(method, endpoint, nil)
		c.Request.Header.Set("Authorization", suite.AdminToken)

		controller.CreateTransaction(c)

		suite.Equal(http.StatusUnauthorized, w.Code)
	}

	// Invalid Payload
	{
		w, c := SetupRouter(controller)
		c.Request, _ = http.NewRequest(method, endpoint, &bytes.Buffer{})
		c.Request.Header.Set("Authorization", suite.TestUserToken)

		controller.CreateTransaction(c)

		suite.Equal(http.StatusBadRequest, w.Code)
	}

	// Failed
	{
		payload := request.DebtorTransactionPayload{
			OTR:          "1000000",
			AssetName:    "Asset",
			TenorLimitID: uuid.New(),
		}
		jsonPayload, _ := json.Marshal(payload)

		w, c := SetupRouter(controller)
		c.Request, _ = http.NewRequest(method, endpoint, bytes.NewBuffer(jsonPayload))
		c.Request.Header.Set("Authorization", suite.TestUserToken)

		controller.CreateTransaction(c)

		var transactionResponse response.DebtorTransactionResponse
		ParseBodyAsStruct(nil, w.Body, &transactionResponse)

		suite.Equal(http.StatusBadRequest, w.Code)
	}
}

func (suite *ControllerTestSuite) TestDebtorController_ListInstallment() {
	controller := controllers.NewDebtorController(suite.Service.DebtorService)
	endpoint := "/debtor/installment/list"
	method := "GET"

	var tenorLimit models.DebtorTenorLimit
	suite.Db.First(&tenorLimit).Where("user_id = ?", suite.TestUser.ID)

	payload := request.DebtorTransactionPayload{
		OTR:          "1000000",
		AssetName:    "Asset",
		TenorLimitID: tenorLimit.ID,
	}
	jsonPayload, _ := json.Marshal(payload)

	_, c := SetupRouter(controller)
	c.Request, _ = http.NewRequest(method, endpoint, bytes.NewBuffer(jsonPayload))
	c.Request.Header.Set("Authorization", suite.TestUserToken)

	controller.CreateTransaction(c)

	// Success
	{
		w, c := SetupRouter(controller)
		c.Request, _ = http.NewRequest(method, endpoint, nil)
		c.Request.Header.Set("Authorization", suite.TestUserToken)

		controller.ListInstallment(c)

		var installmentResponse response.ListDebtorInstallmentResponse
		ParseBodyAsStruct(nil, w.Body, &installmentResponse)

		suite.Equal(http.StatusOK, w.Code)
		suite.Equal(1, len(installmentResponse.Data))
	}

	// Empty Token
	{
		w, c := SetupRouter(controller)
		c.Request, _ = http.NewRequest(method, endpoint, nil)

		controller.ListInstallment(c)

		suite.Equal(http.StatusUnauthorized, w.Code)
	}

	// Unauthorized
	{
		w, c := SetupRouter(controller)
		c.Request, _ = http.NewRequest(method, endpoint, nil)
		c.Request.Header.Set("Authorization", suite.AdminToken)

		controller.ListInstallment(c)

		suite.Equal(http.StatusUnauthorized, w.Code)
	}
}

func (suite *ControllerTestSuite) TestDebtorController_PayInstallmentLine() {
	controller := controllers.NewDebtorController(suite.Service.DebtorService)
	endpoint := "/debtor/installment/pay"
	method := "POST"

	var tenorLimit models.DebtorTenorLimit
	suite.Db.First(&tenorLimit).Where("user_id = ?", suite.TestUser.ID)

	payload := request.DebtorTransactionPayload{
		OTR:          "1000000",
		AssetName:    "Asset",
		TenorLimitID: tenorLimit.ID,
	}
	jsonPayload, _ := json.Marshal(payload)

	_, c := SetupRouter(controller)
	c.Request, _ = http.NewRequest(method, endpoint, bytes.NewBuffer(jsonPayload))
	c.Request.Header.Set("Authorization", suite.TestUserToken)

	controller.CreateTransaction(c)

	var installmentLine models.DebtorInstallmentLine
	suite.Db.First(&installmentLine).Where("user_id = ?", suite.TestUser.ID)

	// Success
	{
		payload := request.DebtorPayInstallmentLinePayload{
			InstallmentLineID: installmentLine.ID,
		}
		jsonPayload, _ := json.Marshal(payload)

		w, c := SetupRouter(controller)
		c.Request, _ = http.NewRequest(method, endpoint, bytes.NewBuffer(jsonPayload))
		c.Request.Header.Set("Authorization", suite.TestUserToken)

		controller.PayInstallmentLine(c)

		var installmentLineResponse response.DebtorInstallmentLineResponse
		ParseBodyAsStruct(nil, w.Body, &installmentLineResponse)

		suite.Equal(http.StatusOK, w.Code)
		suite.Equal(payload.InstallmentLineID.String(), installmentLineResponse.ID)
		suite.Equal("Paid", installmentLineResponse.Status)
	}

	// Empty Token
	{
		w, c := SetupRouter(controller)
		c.Request, _ = http.NewRequest(method, endpoint, nil)

		controller.PayInstallmentLine(c)

		suite.Equal(http.StatusUnauthorized, w.Code)
	}

	// Unauthorized
	{
		w, c := SetupRouter(controller)
		c.Request, _ = http.NewRequest(method, endpoint, nil)
		c.Request.Header.Set("Authorization", suite.AdminToken)

		controller.PayInstallmentLine(c)

		suite.Equal(http.StatusUnauthorized, w.Code)
	}

	// Invalid Payload
	{
		w, c := SetupRouter(controller)
		c.Request, _ = http.NewRequest(method, endpoint, &bytes.Buffer{})
		c.Request.Header.Set("Authorization", suite.TestUserToken)

		controller.PayInstallmentLine(c)

		suite.Equal(http.StatusBadRequest, w.Code)
	}

	// Failed
	{
		payload := request.DebtorPayInstallmentLinePayload{
			InstallmentLineID: uuid.New(),
		}
		jsonPayload, _ := json.Marshal(payload)

		w, c := SetupRouter(controller)
		c.Request, _ = http.NewRequest(method, endpoint, bytes.NewBuffer(jsonPayload))
		c.Request.Header.Set("Authorization", suite.TestUserToken)

		controller.PayInstallmentLine(c)

		suite.Equal(http.StatusBadRequest, w.Code)
	}
}
