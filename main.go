package main

import (
	"credit/config"
	"credit/controllers"
	"credit/services"
	"fmt"

	_ "credit/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Debtor API
// @version         1.0.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	configs, err := config.Load("config.dev.yaml")
	if err != nil {
		panic(err)
	}

	gin.SetMode(configs.App.Mode)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	service := services.NewService(configs.DatabaseConnection)

	apiV1 := router.Group("/api/v1")
	controllers.InitController(service, apiV1)

	router.Run(fmt.Sprintf(":%s", configs.App.Port))
}
