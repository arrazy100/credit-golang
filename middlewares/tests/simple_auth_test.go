package middlewares_tests

import (
	"credit/middlewares"
	"credit/models/enums"
	"credit/utils"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSimpleAuthMiddleware(t *testing.T) {
	router := gin.New()

	router.Use(middlewares.SimpleAuthMiddleware())

	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "Test route")
	})

	userID := uuid.New()
	validToken, _ := utils.GenerateToken(userID, enums.Admin)

	t.Run("ValidToken", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", validToken))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "Test route", w.Body.String())

		c, _ := gin.CreateTestContext(w)
		c.Request = req

		parsed, _ := middlewares.ParseToken(c, true)

		assert.Equal(t, parsed.UserID, userID)
		assert.Equal(t, parsed.Role, enums.Admin)
	})

	t.Run("NoToken", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/test", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("EmptyToken", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", "")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}
