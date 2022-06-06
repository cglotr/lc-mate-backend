package controller_test

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cglotr/lc-mate-backend/controller"
	"github.com/cglotr/lc-mate-backend/leetcode"
	"github.com/cglotr/lc-mate-backend/service"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPostUsersControllerInvalidRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	userServiceMock := service.NewUserServiceMock(ctrl)

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	r.POST("/users", controller.PostUsersController(userServiceMock))

	req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer([]byte("")))
	r.ServeHTTP(w, req)

	assert.Nil(t, err)
	assert.Equal(t, 400, w.Result().StatusCode)

	bytes, err := io.ReadAll(w.Result().Body)

	assert.Nil(t, err)
	assert.Equal(t, `{"error":"unexpected end of JSON input"}`, string(bytes))
}

func TestPostUsersControllerUsersError(t *testing.T) {
	ctrl := gomock.NewController(t)
	userServiceMock := service.NewUserServiceMock(ctrl)

	userServiceMock.EXPECT().
		GetUsers(gomock.Any()).
		Return(nil, errors.New(service.ErrorInvalidUser()))

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	r.POST("/users", controller.PostUsersController(userServiceMock))

	req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer([]byte(`{"usernames":["profile"]}`)))
	r.ServeHTTP(w, req)

	assert.Nil(t, err)
	assert.Equal(t, 400, w.Result().StatusCode)

	bytes, err := io.ReadAll(w.Result().Body)

	assert.Nil(t, err)
	assert.Equal(t, `{"error":"error invalid user"}`, string(bytes))
}

func TestPostUsersControllerHappyPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	userServiceMock := service.NewUserServiceMock(ctrl)

	userServiceMock.EXPECT().
		GetUsers(gomock.Eq([]string{"awice"})).
		Return([]*leetcode.UserInfo{
			{
				Username: "awice",
				Rating:   4000,
				Rank:     "Guardian",
			},
		}, nil)

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	r.POST("/users", controller.PostUsersController(userServiceMock))

	req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer([]byte(`{"usernames":["awice"]}`)))
	r.ServeHTTP(w, req)

	assert.Nil(t, err)
	assert.Equal(t, 200, w.Result().StatusCode)

	bytes, err := io.ReadAll(w.Result().Body)

	assert.Nil(t, err)
	assert.Equal(t, `{"users":[{"username":"awice","rating":4000,"rank":"Guardian"}]}`, string(bytes))
}
