package controller

import (
	"net/http"

	"github.com/cglotr/lc-mate-backend/service"
	"github.com/gin-gonic/gin"
)

func GetUsersController(userService service.UserService) func(c *gin.Context) {
	return func(c *gin.Context) {
		users, err := userService.GetUsers()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	}
}
