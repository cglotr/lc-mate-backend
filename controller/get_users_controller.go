package controller

import (
	"net/http"

	"github.com/cglotr/lc-mate-backend/leetcode"
	"github.com/cglotr/lc-mate-backend/service"
	"github.com/gin-gonic/gin"
)

func GetUsersController(userService service.UserService) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"users": []*leetcode.UserInfo{},
		})
	}
}
