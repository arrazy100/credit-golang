package middlewares

import (
	custom_errors "credit/errors"
	"credit/utils"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	errUnauthorized               = "unauthorized"
	errInvalidAuthorizationHeader = "authorization header is not of type Bearer"
	errTokenNotValid              = "token not valid"
	errInvalidGinContext          = "invalid gin context"
)

func SimpleAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := ParseToken(c, false)
		if err != nil {
			c.JSON(401, gin.H{"errors": custom_errors.Convert(err)})
			c.Abort()
			return
		}

		c.Next()
	}
}

func ParseToken(c *gin.Context, parse bool) (*utils.SimpleAuth, error) {
	if c == nil {
		return nil, errors.New(errInvalidGinContext)
	}

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return nil, errors.New(errUnauthorized)
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return nil, errors.New(errInvalidAuthorizationHeader)
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == "" {
		return nil, errors.New(errUnauthorized)
	}

	data, err := utils.ValidateTokenAndParseData(token)
	if err != nil {
		return nil, errors.New(errTokenNotValid)
	}

	return data, nil
}
