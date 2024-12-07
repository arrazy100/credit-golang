package controllers_tests

import (
	"bytes"
	"credit/controllers"
	"credit/dtos/request"
	"credit/dtos/response"
	"credit/models/enums"
	"credit/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDebtorController_RegisterDebtorSuccess(t *testing.T) {
	mockService := new(MockDebtorService)
	mockService.On("RegisterDebtor", mock.Anything).Return(&response.RegisterDebtorResponse{}, http.StatusOK, nil)

	debtorController := controllers.NewDebtorController(mockService)

	w, c := SetupRouter(debtorController)

	payload := request.RegisterDebtorPayload{}
	jsonPayload, _ := json.Marshal(payload)

	c.Request, _ = http.NewRequest("POST", "/debtor/register", bytes.NewBuffer(jsonPayload))

	validToken, _ := utils.GenerateToken(uuid.New(), enums.Debtor)
	c.Request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", validToken))

	debtorController.RegisterDebtor(c)

	assert.Equal(t, http.StatusOK, w.Code)

	mockService.AssertExpectations(t)
}
