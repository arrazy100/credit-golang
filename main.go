package main

import (
	"credit/config"
	"credit/controllers"
	"credit/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	configs, err := config.Load("config.dev.yaml")
	if err != nil {
		panic(err)
	}

	gin.SetMode(configs.App.Mode)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	service := services.NewService(configs.DatabaseConnection)

	controllers.InitController(service, router)

	router.Run(fmt.Sprintf(":%s", configs.App.Port))
}
