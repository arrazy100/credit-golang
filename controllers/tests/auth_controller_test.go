package controllers_tests

import (
	"bytes"
	"credit/controllers"
	"credit/dtos/request"
	"credit/dtos/response"
	"credit/models/enums"
	"encoding/json"
	"net/http"
)

func (suite *ControllerTestSuite) TestAuthService_Login() {
	controller := controllers.NewAuthController(suite.Service.AuthService)
	endpoint := "/auth/login"
	method := "POST"

	// Success
	{
		payload := request.LoginPayload{
			Email:    "admin@admin.com",
			Password: "password123",
		}
		jsonPayload, _ := json.Marshal(payload)

		w, c := SetupRouter(controller)
		c.Request, _ = http.NewRequest(method, endpoint, bytes.NewBuffer(jsonPayload))

		controller.Login(c)

		var loginResponse response.LoginResponse
		ParseBodyAsStruct(nil, w.Body, &loginResponse)

		suite.Equal(http.StatusOK, w.Code)
		suite.Equal(suite.Admin.ID, loginResponse.User.UserID)
		suite.Equal(suite.Admin.Role.String(), loginResponse.User.Role)
	}

	// Invalid Payload
	{
		w, c := SetupRouter(controller)
		c.Request, _ = http.NewRequest(method, endpoint, &bytes.Buffer{})

		controller.Login(c)

		suite.Equal(http.StatusBadRequest, w.Code)
	}

	// Wrong Password
	{
		payload := request.LoginPayload{
			Email:    "admin@admin.com",
			Password: "password124",
		}
		jsonPayload, _ := json.Marshal(payload)

		w, c := SetupRouter(controller)
		c.Request, _ = http.NewRequest(method, endpoint, bytes.NewBuffer(jsonPayload))

		controller.Login(c)

		suite.Equal(http.StatusUnauthorized, w.Code)
	}
}

func (suite *ControllerTestSuite) TestAuthService_Register() {
	controller := controllers.NewAuthController(suite.Service.AuthService)
	endpoint := "/auth/register"
	method := "POST"

	// Success
	{
		payload := request.RegisterPayload{
			Email:    "newuser@admin.com",
			Password: "password123",
		}
		jsonPayload, _ := json.Marshal(payload)

		w, c := SetupRouter(controller)
		c.Request, _ = http.NewRequest(method, endpoint, bytes.NewBuffer(jsonPayload))

		controller.RegisterUser(c)

		var registerResponse response.RegisterResponse
		ParseBodyAsStruct(nil, w.Body, &registerResponse)

		suite.Equal(http.StatusCreated, w.Code)
		suite.Equal(payload.Email, registerResponse.Email)
		suite.Equal(enums.Debtor.String(), registerResponse.Role)
	}

	// Invalid Payload
	{
		w, c := SetupRouter(controller)
		c.Request, _ = http.NewRequest(method, endpoint, &bytes.Buffer{})

		controller.RegisterUser(c)

		suite.Equal(http.StatusBadRequest, w.Code)
	}

	// Already Registered
	{
		payload := request.RegisterPayload{
			Email:    "admin@admin.com",
			Password: "password123",
		}
		jsonPayload, _ := json.Marshal(payload)

		w, c := SetupRouter(controller)
		c.Request, _ = http.NewRequest(method, endpoint, bytes.NewBuffer(jsonPayload))

		controller.RegisterUser(c)

		suite.Equal(http.StatusBadRequest, w.Code)
	}
}
