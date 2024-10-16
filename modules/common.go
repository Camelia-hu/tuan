package modules

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name       string
	StudentNum string
	PassWord   string
	Salt       string
}

type RegisterReq struct {
	StudentNum string `json:"studentNum"`
	PassWord   string `json:"passWord"`
}

type RegisterResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	User User   `json:"user"`
}

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
