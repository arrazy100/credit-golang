package controllers

import (
	"credit/controllers/interfaces"
	"credit/services"

	"github.com/gin-gonic/gin"
)

func InitController(service *services.Service, router *gin.Engine) {
	controllers := []interfaces.IController{
		NewAuthController(service.AuthService),
	}

	for _, controller := range controllers {
		controller.SetupGroup(router)
	}
}
