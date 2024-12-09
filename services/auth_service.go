package services

import (
	"credit/dtos/request"
	"credit/dtos/response"
	"credit/models/base"
	"credit/models/enums"
	"credit/services/interfaces"
	"credit/utils"
	validations "credit/validations"
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

func (s *AuthService) Login(payload request.LoginPayload) (*response.LoginResponse, int, *validations.ErrorValidation) {
	validationError := validations.ValidateStruct(payload)
	if validationError != nil {
		return nil, statusBadRequest, validationError
	}

	var userData base.User
	result := s.db.Where("email = ?", payload.Email).First(&userData)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, statusBadRequest, validations.Convert(fmt.Errorf(errAccountNotFound, payload.Email))
		}
		return nil, statusServerError, validations.Convert(result.Error)
	}

	if !utils.CheckPasswordHash(payload.Password, userData.Password) {
		return nil, statusUnauthorized, validations.Convert(errors.New(errInvalidPassword))
	}

	token, err := utils.GenerateToken(userData.ID, userData.Role)
	if err != nil {
		return nil, statusServerError, validations.Convert(err)
	}

	response := &response.LoginResponse{
		Token: token,
		User: response.UserResponse{
			UserID: userData.ID,
			Role:   userData.Role.String(),
		},
	}

	return response, statusSuccess, nil
}

func (s *AuthService) RegisterUser(payload request.RegisterPayload) (*response.RegisterResponse, int, *validations.ErrorValidation) {
	validationError := validations.ValidateStruct(payload)
	if validationError != nil {
		return nil, statusBadRequest, validationError
	}

	var userData base.User
	if err := s.db.Where("email = ?", payload.Email).First(&userData).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, statusServerError, validations.Convert(err)
	}

	if userData.ID != uuid.Nil {
		return nil, statusBadRequest, validations.Convert(fmt.Errorf("email %s already registered", payload.Email))
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		return nil, statusServerError, validations.Convert(err)
	}

	newUser := base.User{
		ID:       uuid.New(),
		Email:    payload.Email,
		Password: hashedPassword,
		Role:     enums.Debtor,
	}

	if err := s.db.Create(&newUser).Error; err != nil {
		return nil, statusServerError, validations.Convert(err)
	}

	response := &response.RegisterResponse{
		UserID: newUser.ID,
		Email:  newUser.Email,
		Role:   newUser.Role.String(),
	}

	return response, statusCreated, nil
}
