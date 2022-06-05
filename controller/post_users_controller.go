package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/cglotr/lc-mate-backend/service"
	"github.com/gin-gonic/gin"
)

func PostUsersController(userService service.UserService) func(c *gin.Context) {
	return func(c *gin.Context) {
		bytes, _ := ioutil.ReadAll(c.Request.Body)
		type PostUsersRequestJson struct {
			Usernames []string `json:"usernames"`
		}
		var postUsersRequestJson PostUsersRequestJson
		err := json.Unmarshal(bytes, &postUsersRequestJson)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		users, err := userService.GetUsers(postUsersRequestJson.Usernames)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	}
}
