package services

import (
	"credit/services/interfaces"

	"gorm.io/gorm"
)

type Service struct {
	AuthService interfaces.AuthInterface
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		AuthService: NewAuthService(db),
	}
}
