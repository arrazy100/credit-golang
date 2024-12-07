package controllers_tests

import (
	"bytes"
	"credit/controllers/interfaces"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func ParseBodyAsStruct[T any](t *testing.T, body *bytes.Buffer, result *T) {
	var data T
	if err := json.Unmarshal(body.Bytes(), &data); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	dataJSON, _ := json.Marshal(data)
	if err := json.Unmarshal(dataJSON, result); err != nil {
		t.Fatalf("Failed to unmarshal 'data' field: %v", err)
	}
}

func SetupRouter(controller interfaces.IController) (*httptest.ResponseRecorder, *gin.Context) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	group := r.Group("")

	controller.SetupGroup(group)

	return w, c
}
