package controllers

import (
	"credit/controllers/interfaces"
	"credit/middlewares"
	"credit/models/enums"
	service_interface "credit/services/interfaces"
	validations "credit/validations"
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
// @Failure      400  {object}  validations.ErrorValidation
// @Failure      500  {object}  validations.ErrorValidation
// @Router       /list/debtor [get]
func (c *AdminController) ListDebtor(ctx *gin.Context) {
	auth, err := middlewares.ParseToken(ctx, true)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"errors": validations.Convert(err)})
		return
	}

	if err := auth.ValidateRole(enums.Admin); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"errors": validations.Convert(err)})
		return
	}

	response, status, validationErrors := c.AdminService.ListDebtor()
	if validationErrors != nil {
		ctx.JSON(status, gin.H{"errors": validationErrors})
		return
	}

	ctx.JSON(status, response)
}
