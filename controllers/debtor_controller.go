package controllers

import (
	"credit/controllers/interfaces"
	"credit/dtos/request"
	custom_errors "credit/errors"
	"credit/middlewares"
	"credit/models/enums"
	service_interface "credit/services/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

var _ interfaces.IController = (*DebtorController)(nil)

type DebtorController struct {
	DebtorService service_interface.DebtorInterface
}

func NewDebtorController(debtorService service_interface.DebtorInterface) *DebtorController {
	return &DebtorController{
		DebtorService: debtorService,
	}
}

func (c *DebtorController) SetupGroup(router *gin.RouterGroup) {
	group := router.Group("/debtor")
	group.Use(middlewares.SimpleAuthMiddleware())

	group.POST("/register", c.RegisterDebtor)
	group.GET("/detail", c.DetailDebtor)
}

// Register Debtor godoc
// @Summary      Register Debtor
// @Description  Register Debtor to get Tenor Limits
// @Tags         debtor
// @Accept       json
// @Produce      json
// @Param        request body request.RegisterDebtorPayload true "Body"
// @Success      200  {object}  response.RegisterDebtorResponse
// @Failure      400  {object}  custom_errors.ErrorValidation
// @Failure      500  {object}  custom_errors.ErrorValidation
// @Router       /debtor/register [post]
func (c *DebtorController) RegisterDebtor(ctx *gin.Context) {
	auth, err := middlewares.ParseToken(ctx, true)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"errors": custom_errors.Convert(err)})
	}

	if err := auth.ValidateRole(enums.Debtor); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"errors": custom_errors.Convert(err)})
	}

	var payload request.RegisterDebtorPayload
	if err := ctx.ShouldBindBodyWithJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": custom_errors.Convert(err)})
		return
	}

	response, status, validationErrors := c.DebtorService.RegisterDebtor(auth.UserID, payload)
	if validationErrors != nil {
		ctx.JSON(status, gin.H{"errors": validationErrors})
		return
	}

	ctx.JSON(status, response)
}

// Detail Debtor godoc
// @Summary      Detail Debtor
// @Description  Detail Debtor of logged in User
// @Tags         debtor
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.DebtorResponse
// @Failure      400  {object}  custom_errors.ErrorValidation
// @Failure      500  {object}  custom_errors.ErrorValidation
// @Router       /debtor/detail [get]
func (c *DebtorController) DetailDebtor(ctx *gin.Context) {
	auth, err := middlewares.ParseToken(ctx, true)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"errors": custom_errors.Convert(err)})
	}

	if err := auth.ValidateRole(enums.Debtor); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"errors": custom_errors.Convert(err)})
	}

	response, status, validationErrors := c.DebtorService.DetailDebtor(auth.UserID)
	if validationErrors != nil {
		ctx.JSON(status, gin.H{"errors": validationErrors})
	}

	ctx.JSON(status, response)
}
