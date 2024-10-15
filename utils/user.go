package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/Camelia-hu/tuan/dao"
	"github.com/Camelia-hu/tuan/modules"
	"golang.org/x/exp/rand"
	"gorm.io/gorm"
	"log"
	"time"
)

func ExistOrNot(studentnum string) bool {
	var user modules.User
	err := dao.DB.Where("student_num = ?", studentnum).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("用户不存在，可以创建")
		return false
	}
	return true
}

func CreateUser(studentnum string, password string) {
	user := modules.User{
		Model:      gorm.Model{},
		Name:       "",
		StudentNum: studentnum,
		PassWord:   password,
	}
	dao.DB.Create(&user)
}

func GenerateSalt() string {
	rand.Seed(uint64(time.Now().UnixNano()))
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	salt := make([]byte, 16)
	for i := range salt {
		salt[i] = letters[rand.Intn(len(letters))]
	}
	return string(salt)
}

func HashPassword(password string, salt string) string {
	h := sha256.New()
	h.Write([]byte(password + salt))
	return hex.EncodeToString(h.Sum(nil))
}
