package service

import (
	"github.com/Camelia-hu/tuan/modules"
	"github.com/Camelia-hu/tuan/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"time"
)

func Register(c *gin.Context) {
	var registerreq modules.RegisterReq
	err := c.ShouldBindJSON(&registerreq)
	if err != nil {
		log.Println(err)
	}
	//用户名或密码为空
	if registerreq.StudentNum == "" || registerreq.PassWord == "" {
		c.JSON(400, modules.Response{
			Code: 1,
			Msg:  "请输入用户名或密码",
		})
		return
	}
	//检查用户名是否已存在
	if utils.ExistOrNot(registerreq.StudentNum) {
		c.JSON(400, modules.Response{
			Code: 1,
			Msg:  "用户名已存在",
		})
		return
	}
	//对密码进行加密
	salt := utils.GenerateSalt()
	hashPassword := utils.HashPassword(registerreq.PassWord, salt)
	if err != nil {
		log.Println(err)
	}
	//创建用户
	utils.CreateUser(registerreq.StudentNum, hashPassword)
	c.JSON(200, modules.RegisterResponse{
		Code: 0,
		Msg:  "注册成功",
		User: modules.User{
			Model: gorm.Model{
				ID:        0,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: gorm.DeletedAt{},
			},
			StudentNum: registerreq.StudentNum,
			PassWord:   hashPassword,
		},
	})
}

func Login(c *gin.Context) {

}
