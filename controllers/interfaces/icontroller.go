package interfaces

import "github.com/gin-gonic/gin"

type IController interface {
	SetupGroup(router *gin.Engine)
}
