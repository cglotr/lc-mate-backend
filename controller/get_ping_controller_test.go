package controller_test

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/cglotr/lc-mate-backend/controller"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetPingController(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	controller.GetPingController()(c)
	bytes, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, `{"message":"pong"}`, string(bytes))
}
