package middlewares_tests

import (
	"credit/middlewares"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSimpleAuthMiddleware(t *testing.T) {
	router := gin.New()

	router.Use(middlewares.SimpleAuthMiddleware())

	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "Test route")
	})

	t.Run("ValidToken", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", "Bearer valid_token")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "Test route", w.Body.String())
	})

	t.Run("NoToken", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/test", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Equal(t, `{"error":"Unauthorized"}`, w.Body.String())
	})

	t.Run("EmptyToken", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", "")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Equal(t, `{"error":"Unauthorized"}`, w.Body.String())
	})
}
