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
	json.Unmarshal(body.Bytes(), &data)

	dataJSON, _ := json.Marshal(data)
	json.Unmarshal(dataJSON, result)
}

func SetupRouter(controller interfaces.IController) (*httptest.ResponseRecorder, *gin.Context) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)

	group := r.Group("")

	controller.SetupGroup(group)

	return w, c
}
