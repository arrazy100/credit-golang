package services

import (
	"credit/services/interfaces"
	"net/http"

	"gorm.io/gorm"
)

const (
	statusSuccess      = http.StatusOK
	statusCreated      = http.StatusCreated
	statusBadRequest   = http.StatusBadRequest
	statusUnauthorized = http.StatusUnauthorized
	statusServerError  = http.StatusInternalServerError
)

type Service struct {
	AuthService   interfaces.AuthInterface
	DebtorService interfaces.DebtorInterface
	AdminService  interfaces.AdminInterface
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		AuthService:   NewAuthService(db),
		DebtorService: NewDebtorService(db),
		AdminService:  NewAdminService(db),
	}
}
