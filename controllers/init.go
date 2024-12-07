package controllers

import (
	"credit/controllers/interfaces"
	"credit/services"

	"github.com/gin-gonic/gin"
)

func InitController(service *services.Service, router *gin.RouterGroup) {
	controllers := []interfaces.IController{
		NewAuthController(service.AuthService),
		NewDebtorController(service.DebtorService),
		NewAdminController(service.AdminService),
	}

	for _, controller := range controllers {
		controller.SetupGroup(router)
	}
}
