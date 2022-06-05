package main

import (
	"github.com/cglotr/lc-mate-backend/controller"
	"github.com/cglotr/lc-mate-backend/service"
	"github.com/gin-gonic/gin"
)

func routes(r *gin.Engine, userService service.UserService) {
	r.GET("/ping", controller.GetPingController())
	r.POST("/users", controller.PostUsersController(userService))
}
