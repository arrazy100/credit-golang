package controllers

import (
	"credit/controllers/interfaces"
	"credit/dtos/request"
	custom_errors "credit/errors"
	service_interface "credit/services/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

var _ interfaces.IController = (*AuthController)(nil)

type AuthController struct {
	AuthService service_interface.AuthInterface
}

func NewAuthController(authService service_interface.AuthInterface) *AuthController {
	return &AuthController{
		AuthService: authService,
	}
}

func (c *AuthController) SetupGroup(router *gin.RouterGroup) {
	group := router.Group("/auth")

	group.POST("/login", c.Login)
	group.POST("/register", c.RegisterUser)
}

// Login godoc
// @Summary      Login
// @Description  Login to get token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body request.LoginPayload true "Body"
// @Success      200  {object}  response.LoginResponse
// @Failure      400  {object}  custom_errors.ErrorValidation
// @Failure      500  {object}  custom_errors.ErrorValidation
// @Router       /auth/login [post]
func (c *AuthController) Login(ctx *gin.Context) {
	var payload request.LoginPayload
	if err := ctx.ShouldBindBodyWithJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, custom_errors.Convert(err))
		return
	}

	response, status, validationErrors := c.AuthService.Login(payload)
	if validationErrors != nil {
		ctx.JSON(status, gin.H{"errors": validationErrors})
		return
	}

	ctx.JSON(status, response)
}

// Register User godoc
// @Summary      Register
// @Description  Register Debtor User
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body request.RegisterPayload true "Body"
// @Success      200  {object}  response.RegisterResponse
// @Failure      400  {object}  custom_errors.ErrorValidation
// @Failure      500  {object}  custom_errors.ErrorValidation
// @Router       /auth/register [post]
func (c *AuthController) RegisterUser(ctx *gin.Context) {
	var payload request.RegisterPayload
	if err := ctx.ShouldBindBodyWithJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": custom_errors.Convert(err)})
		return
	}

	response, status, validationErrors := c.AuthService.RegisterUser(payload)
	if validationErrors != nil {
		ctx.JSON(status, gin.H{"errors": validationErrors})
		return
	}

	ctx.JSON(status, response)
}
