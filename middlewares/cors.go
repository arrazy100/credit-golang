package middlewares

import (
	"credit/config"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware(config config.Server) gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:  config.AllowOrigins,
		AllowMethods:  config.AllowMethods,
		AllowHeaders:  config.AllowHeaders,
		ExposeHeaders: config.ExposeHeaders,
		MaxAge:        12 * time.Hour,
	})
}
