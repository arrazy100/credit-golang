package middlewares

import (
	"credit/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	errUnsupportedContentType = "unsupported Media Type"
)

func ContentTypeMiddleware(config config.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		contentType := c.GetHeader("Content-Type")

		allowed := false
		for _, allowedType := range config.AllowContentType {
			if contentType == allowedType {
				allowed = true
				break
			}
		}

		if !allowed {
			c.AbortWithStatusJSON(http.StatusUnsupportedMediaType, gin.H{
				"error": errUnsupportedContentType,
			})
			return
		}

		c.Next()
	}
}
