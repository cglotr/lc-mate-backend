package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPingController() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}
}
