package controllers

import (
	"credit/controllers/interfaces"
	"credit/dtos/request"
	"credit/middlewares"
	"credit/models/enums"
	service_interface "credit/services/interfaces"
	validations "credit/validations"
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

	group.POST("/register", c.Register)
	group.GET("/detail", c.Detail)
	group.POST("/transaction", c.CreateTransaction)
	group.GET("/installment/list", c.ListInstallment)
	group.POST("/installment/pay", c.PayInstallmentLine)
}

// Register Debtor godoc
// @Summary      Register Debtor
// @Description  Register Debtor to get Tenor Limits
// @Tags         debtor
// @Accept       json
// @Produce      json
// @Param        request body request.RegisterDebtorPayload true "Body"
// @Success      200  {object}  response.RegisterDebtorResponse
// @Failure      400  {object}  validations.ErrorValidation
// @Failure      500  {object}  validations.ErrorValidation
// @Router       /debtor/register [post]
func (c *DebtorController) Register(ctx *gin.Context) {
	auth, err := middlewares.ParseToken(ctx, true)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"errors": validations.Convert(err)})
		return
	}

	if err := auth.ValidateRole(enums.Debtor); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"errors": validations.Convert(err)})
		return
	}

	var payload request.RegisterDebtorPayload
	if err := ctx.ShouldBindBodyWithJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": validations.Convert(err)})
		return
	}

	response, status, validationErrors := c.DebtorService.Register(auth.UserID, payload)
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
// @Failure      400  {object}  validations.ErrorValidation
// @Failure      500  {object}  validations.ErrorValidation
// @Router       /debtor/detail [get]
func (c *DebtorController) Detail(ctx *gin.Context) {
	auth, err := middlewares.ParseToken(ctx, true)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"errors": validations.Convert(err)})
		return
	}

	if err := auth.ValidateRole(enums.Debtor); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"errors": validations.Convert(err)})
		return
	}

	response, status, validationErrors := c.DebtorService.Detail(auth.UserID)
	if validationErrors != nil {
		ctx.JSON(status, gin.H{"errors": validationErrors})
		return
	}

	ctx.JSON(status, response)
}

// Create Debtor Transaction godoc
// @Summary      Create Debtor Transaction
// @Description  Create a new Debtor Transaction
// @Tags         debtor
// @Accept       json
// @Produce      json
// @Param        request body request.DebtorTransactionPayload true "Body"
// @Success      200  {object}  response.DebtorTransactionResponse
// @Failure      400  {object}  validations.ErrorValidation
// @Failure      500  {object}  validations.ErrorValidation
// @Router       /debtor/transaction [post]
func (c *DebtorController) CreateTransaction(ctx *gin.Context) {
	auth, err := middlewares.ParseToken(ctx, true)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"errors": validations.Convert(err)})
		return
	}

	if err := auth.ValidateRole(enums.Debtor); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"errors": validations.Convert(err)})
		return
	}

	var payload request.DebtorTransactionPayload
	if err := ctx.ShouldBindBodyWithJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": validations.Convert(err)})
		return
	}

	response, status, validationErrors := c.DebtorService.CreateTransaction(auth.UserID, payload)
	if validationErrors != nil {
		ctx.JSON(status, gin.H{"errors": validationErrors})
		return
	}

	ctx.JSON(status, response)
}

// List Debtor Installment godoc
// @Summary      List Debtor Installment
// @Description  List Debtor Installment of logged in User
// @Tags         debtor
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.ListDebtorInstallmentResponse
// @Failure      400  {object}  validations.ErrorValidation
// @Failure      500  {object}  validations.ErrorValidation
// @Router       /debtor/installment/list [get]
func (c *DebtorController) ListInstallment(ctx *gin.Context) {
	auth, err := middlewares.ParseToken(ctx, true)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"errors": validations.Convert(err)})
		return
	}

	if err := auth.ValidateRole(enums.Debtor); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"errors": validations.Convert(err)})
		return
	}

	response, status, validationErrors := c.DebtorService.ListInstallment(auth.UserID)
	if validationErrors != nil {
		ctx.JSON(status, gin.H{"errors": validationErrors})
		return
	}

	ctx.JSON(status, response)
}

// Pay Debtor Installment Line godoc
// @Summary      Pay Debtor Installment Line
// @Description  Pay Debtor Installment Line of logged in User
// @Tags         debtor
// @Accept       json
// @Produce      json
// @Param        request body request.DebtorPayInstallmentLinePayload true "Body"
// @Success      200  {object}  response.DebtorInstallmentLineResponse
// @Failure      400  {object}  validations.ErrorValidation
// @Failure      500  {object}  validations.ErrorValidation
// @Router       /debtor/installment/pay [post]
func (c *DebtorController) PayInstallmentLine(ctx *gin.Context) {
	auth, err := middlewares.ParseToken(ctx, true)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"errors": validations.Convert(err)})
		return
	}

	if err := auth.ValidateRole(enums.Debtor); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"errors": validations.Convert(err)})
		return
	}

	var payload request.DebtorPayInstallmentLinePayload
	if err := ctx.ShouldBindBodyWithJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": validations.Convert(err)})
		return
	}

	response, status, validationErrors := c.DebtorService.PayInstallmentLine(auth.UserID, payload)
	if validationErrors != nil {
		ctx.JSON(status, gin.H{"errors": validationErrors})
		return
	}

	ctx.JSON(status, response)
}
