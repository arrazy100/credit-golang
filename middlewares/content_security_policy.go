package middlewares

import (
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
)

func ContentSecurityPolicyMiddleware() gin.HandlerFunc {
	return secure.New(secure.Config{
		ContentSecurityPolicy: "default-src 'self'; script-src 'self' https://cdnjs.cloudflare.com https://unpkg.com; style-src 'self' https://cdnjs.cloudflare.com https://unpkg.com; img-src 'self' https://swagger.io; font-src 'self' https://cdnjs.cloudflare.com; connect-src 'self';",
	})
}
