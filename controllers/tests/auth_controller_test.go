package controllers_tests

import (
	"bytes"
	"credit/controllers"
	"credit/dtos/request"
	"credit/dtos/response"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) Login(email, password string) (*response.LoginResponse, error) {
	args := m.Called(email, password)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*response.LoginResponse), args.Error(1)
}

func TestAuthController_Login(t *testing.T) {
	router := gin.Default()

	mockAuthService := new(MockAuthService)

	authController := controllers.NewAuthController(mockAuthService)
	authController.SetupGroup(router)

	testCases := []struct {
		name           string
		payload        interface{}
		expectedStatus int
		expectedError  string
		mockService    func(service *MockAuthService)
	}{
		{
			name: "ValidLogin",
			payload: request.LoginPayload{
				Email:    "test@example.com",
				Password: "password",
			},
			expectedStatus: http.StatusOK,
			expectedError:  "",
			mockService: func(service *MockAuthService) {
				service.On("Login", "test@example.com", "password").Return(&response.LoginResponse{
					Token: "test_token",
				}, nil)
			},
		},
		{
			name: "InvalidLogin",
			payload: request.LoginPayload{
				Email:    "test@example.com",
				Password: "wrong_password",
			},
			expectedStatus: http.StatusUnauthorized,
			expectedError:  "invalid credentials",
			mockService: func(service *MockAuthService) {
				service.On("Login", "test@example.com", "wrong_password").Return(nil, errors.New("invalid credentials"))
			},
		},
		{
			name:           "InvalidPayload",
			payload:        "invalid_payload",
			expectedStatus: http.StatusBadRequest,
			expectedError:  "json: cannot unmarshal string into Go value of type request.LoginPayload",
			mockService:    func(service *MockAuthService) {},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockService(mockAuthService)

			payloadBytes, _ := json.Marshal(tc.payload)
			req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(payloadBytes))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)

			var respBody map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &respBody)
			assert.NoError(t, err)

			if tc.expectedError != "" {
				assert.Contains(t, respBody["error"], tc.expectedError)
			} else {
				assert.Equal(t, "test_token", respBody["data"].(map[string]interface{})["token"])
			}

			mockAuthService.AssertExpectations(t)
		})
	}
}
