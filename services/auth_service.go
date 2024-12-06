package services

import (
	"credit/dtos/response"
	"credit/models/base"
	"credit/services/interfaces"
	"credit/utils"
	"errors"

	"gorm.io/gorm"
)

var _ interfaces.AuthInterface = (*AuthService)(nil)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{db: db}
}

func (s *AuthService) Login(email, password string) (*response.LoginResponse, error) {
	var userData base.User
	result := s.db.Where("email = ?", email).First(&userData)
	if result.Error != nil {
		return nil, result.Error
	}

	if !utils.CheckPasswordHash(password, userData.Password) {
		return nil, errors.New("invalid password")
	}

	token, err := utils.GenerateToken(userData.ID, userData.Role)
	if err != nil {
		return nil, err
	}

	response := &response.LoginResponse{
		Token: token,
		User: utils.SimpleAuth{
			UserID: userData.ID,
			Role:   userData.Role,
		},
	}

	return response, nil
}
