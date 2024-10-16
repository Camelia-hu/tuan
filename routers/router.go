package routers

import (
	"github.com/Camelia-hu/tuan/service"
	"github.com/Camelia-hu/tuan/utils"
	"github.com/gin-gonic/gin"
)

func RoutersInit() {
	r := gin.Default()

	user := r.Group("/user")
	{
		user.POST("/register", service.Register)
		user.GET("/login", service.Login)
		auth := user.Group("/auth", utils.AuthToken())
		{
			auth.POST("/upload", service.UploadUser)
		}
	}

	activity := r.Group("/activity")
	{
		activity.POST("/upload", service.Upload)
	}

	r.Run()
}
