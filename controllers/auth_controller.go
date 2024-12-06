package controllers

import (
	"credit/controllers/interfaces"
	"credit/dtos/request"
	service_interface "credit/services/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService service_interface.AuthInterface
}

var _ interfaces.IController = (*AuthController)(nil)

func NewAuthController(authService service_interface.AuthInterface) *AuthController {
	return &AuthController{
		AuthService: authService,
	}
}

func (c *AuthController) SetupGroup(router *gin.Engine) {
	group := router.Group("/auth")

	group.POST("/login", c.Login)
}

func (c *AuthController) Login(ctx *gin.Context) {
	var payload request.LoginPayload
	if err := ctx.ShouldBindBodyWithJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := c.AuthService.Login(payload.Email, payload.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": response})
}
