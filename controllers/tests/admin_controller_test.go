package controllers_tests

import (
	"credit/controllers"
	"credit/dtos/response"
	"net/http"
)

func (suite *ControllerTestSuite) TestAdminService_ListDebtor() {
	controller := controllers.NewAdminController(suite.Service.AdminService)
	endpoint := "/admin/list/debtor"
	method := "GET"

	// Success
	{
		w, c := SetupRouter(controller)
		c.Request, _ = http.NewRequest(method, endpoint, nil)
		c.Request.Header.Set("Authorization", suite.AdminToken)

		controller.ListDebtor(c)

		var listDebtorResponse response.ListDebtorResponse
		ParseBodyAsStruct(nil, w.Body, &listDebtorResponse)

		suite.Equal(http.StatusOK, w.Code)
		suite.Equal(1, len(listDebtorResponse.Data))
	}

	// Empty Token
	{
		w, c := SetupRouter(controller)
		c.Request, _ = http.NewRequest(method, endpoint, nil)

		controller.ListDebtor(c)

		suite.Equal(http.StatusUnauthorized, w.Code)
	}

	// Unauthorized
	{
		w, c := SetupRouter(controller)
		c.Request, _ = http.NewRequest(method, endpoint, nil)
		c.Request.Header.Set("Authorization", suite.DebtorToken)

		controller.ListDebtor(c)

		suite.Equal(http.StatusUnauthorized, w.Code)
	}
}
