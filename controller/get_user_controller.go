package controller

import (
	"net/http"

	"github.com/cglotr/lc-mate-backend/service"
	"github.com/gin-gonic/gin"
)

func GetUserController(userService service.UserService) func(c *gin.Context) {
	return func(c *gin.Context) {
		username := c.Request.URL.Query().Get("username")
		userInfo, err := userService.GetUser(username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"user": userInfo,
		})
	}
}
