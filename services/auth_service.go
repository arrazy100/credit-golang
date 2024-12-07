package services

import (
	"credit/dtos/request"
	"credit/dtos/response"
	custom_errors "credit/errors"
	"credit/models/base"
	"credit/models/enums"
	"credit/services/interfaces"
	"credit/utils"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var _ interfaces.AuthInterface = (*AuthService)(nil)

const (
	errInvalidPassword = "invalid password"
	errAccountNotFound = "account with email %s not found"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{db: db}
}

func (s *AuthService) Login(payload request.LoginPayload) (*response.LoginResponse, int, *custom_errors.ErrorValidation) {
	validationError := custom_errors.ValidateStruct(payload)
	if validationError != nil {
		return nil, statusBadRequest, validationError
	}

	var userData base.User
	result := s.db.Where("email = ?", payload.Email).First(&userData)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, statusBadRequest, custom_errors.Convert(fmt.Errorf(errAccountNotFound, payload.Email))
		}
		return nil, statusServerError, custom_errors.Convert(result.Error)
	}

	if !utils.CheckPasswordHash(payload.Password, userData.Password) {
		return nil, statusUnauthorized, custom_errors.Convert(errors.New(errInvalidPassword))
	}

	token, err := utils.GenerateToken(userData.ID, userData.Role)
	if err != nil {
		return nil, statusServerError, custom_errors.Convert(err)
	}

	response := &response.LoginResponse{
		Token: token,
		User: utils.SimpleAuth{
			UserID: userData.ID,
			Role:   userData.Role,
		},
	}

	return response, statusSuccess, nil
}

func (s *AuthService) RegisterUser(payload request.RegisterPayload) (*response.RegisterResponse, int, *custom_errors.ErrorValidation) {
	validationError := custom_errors.ValidateStruct(payload)
	if validationError != nil {
		return nil, statusBadRequest, validationError
	}

	var userData base.User
	result := s.db.Where("email = ?", payload.Email).First(&userData)
	if result.Error == nil {
		return nil, statusBadRequest, custom_errors.Convert(fmt.Errorf("email %s already registered", payload.Email))
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		return nil, statusServerError, custom_errors.Convert(err)
	}

	newUser := base.User{
		ID:       uuid.New(),
		Email:    payload.Email,
		Password: hashedPassword,
		Role:     enums.Debtor,
	}

	result = s.db.Create(&newUser)
	if result.Error != nil {
		return nil, statusServerError, custom_errors.Convert(result.Error)
	}

	response := &response.RegisterResponse{
		UserID: newUser.ID,
		Email:  newUser.Email,
		Role:   newUser.Role,
	}

	return response, statusCreated, nil
}
