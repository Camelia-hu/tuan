package routers

import (
	"github.com/Camelia-hu/tuan/service"
	"github.com/gin-gonic/gin"
)

func RoutersInit() {
	r := gin.Default()

	user := r.Group("/user")
	{
		user.POST("/register", service.Register)
		user.GET("/login", service.Login)
	}

	activity := r.Group("/activity")
	{
		activity.POST("/upload", service.Upload)
	}

	r.Run()
}
