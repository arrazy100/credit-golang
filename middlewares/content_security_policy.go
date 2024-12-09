package middlewares

import (
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
)

func ContentSecurityPolicyMiddleware() gin.HandlerFunc {
	return secure.New(secure.Config{
		ContentSecurityPolicy: "default-src 'self'; script-src 'self'; style-src 'self';",
	})
}
