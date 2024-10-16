package service

import (
	"errors"
	"github.com/Camelia-hu/tuan/dao"
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
	utils.CreateUser(registerreq.StudentNum, hashPassword, salt)
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
	var user modules.User
	studentNum := c.Query("studentNum")
	passWord := c.Query("passWord")

	//验证输入是否合法
	if studentNum == "" || passWord == "" {
		c.JSON(400, modules.Response{
			Code: 1,
			Msg:  "请输入用户名或密码",
		})
		return
	}

	//用户名不存在
	err := dao.DB.Where("student_num = ?", studentNum).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, modules.Response{
			Code: 1,
			Msg:  "该用户名不存在",
		})
		return
	}

	//密码错误
	hashPassWord := utils.HashPassword(passWord, user.Salt)
	if hashPassWord != user.PassWord {
		c.JSON(400, modules.Response{
			Code: 1,
			Msg:  "密码错误",
		})
		return
	}

	//token生成
	token, err := utils.JwtGenerate(int(user.ID), user.Name)
	if err != nil {
		log.Println(err)
		c.JSON(400, modules.Response{
			Code: 1,
			Msg:  "token生成失败",
		})
		return
	}

	//登陆成功
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "登陆成功",
		"user": modules.User{
			Model:      gorm.Model{},
			Name:       user.Name,
			StudentNum: user.StudentNum,
		},
		"token": token,
	})
}

func UploadUser(c *gin.Context) {
	type UploadUser struct {
		Name       string `form:"name" json:"name"`
		StudentNum string `form:"studentNum" json:"studentNum"`
	}
	uploadUser := UploadUser{}
	err := c.ShouldBind(&uploadUser)
	if err != nil {
		log.Println("Bind err : ", err)
		c.JSON(400, modules.Response{
			Code: 1,
			Msg:  "参数绑定错误",
		})
		return
	}
	err = dao.DB.Model(&modules.User{}).Where("student_num = ?", uploadUser.StudentNum).Update("name", uploadUser.Name).Error
	if err != nil {
		log.Println("upload user err : ", err)
		c.JSON(400, modules.Response{
			Code: 1,
			Msg:  "数据库操作失败",
		})
		return
	}
	c.JSON(200, modules.Response{
		Code: 0,
		Msg:  "用户信息添加成功",
	})
}
