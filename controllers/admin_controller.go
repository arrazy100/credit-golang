package controllers

import (
	"credit/controllers/interfaces"
	custom_errors "credit/errors"
	"credit/middlewares"
	"credit/models/enums"
	service_interface "credit/services/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

var _ interfaces.IController = (*AdminController)(nil)

type AdminController struct {
	AdminService service_interface.AdminInterface
}

func NewAdminController(adminService service_interface.AdminInterface) *AdminController {
	return &AdminController{
		AdminService: adminService,
	}
}

func (c *AdminController) SetupGroup(router *gin.RouterGroup) {
	group := router.Group("/admin")
	group.Use(middlewares.SimpleAuthMiddleware())

	group.POST("/list/debtor", c.ListDebtor)
}

// List Debtor godoc
// @Summary      List Debtor
// @Description  List all registered Debtor
// @Tags         admin
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.ListDebtorResponse
// @Failure      400  {object}  custom_errors.ErrorValidation
// @Failure      500  {object}  custom_errors.ErrorValidation
// @Router       /list/debtor [get]
func (c *AdminController) ListDebtor(ctx *gin.Context) {
	auth, err := middlewares.ParseToken(ctx, true)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"errors": custom_errors.Convert(err)})
	}

	if err := auth.ValidateRole(enums.Admin); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"errors": custom_errors.Convert(err)})
	}

	response, status, validationErrors := c.AdminService.ListDebtor()
	if validationErrors != nil {
		ctx.JSON(status, gin.H{"errors": validationErrors})
		return
	}

	ctx.JSON(status, response)
}
