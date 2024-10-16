package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/Camelia-hu/tuan/config"
	"github.com/Camelia-hu/tuan/dao"
	"github.com/Camelia-hu/tuan/modules"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/exp/rand"
	"gorm.io/gorm"
	"log"
	"time"
)

type JwtClaims struct {
	Id   int
	Name string
	jwt.RegisteredClaims
}

func ExistOrNot(studentnum string) bool {
	var user modules.User
	err := dao.DB.Where("student_num = ?", studentnum).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("用户不存在，可以创建")
		return false
	}
	return true
}

func CreateUser(studentnum string, password string, salt string) {
	user := modules.User{
		Model:      gorm.Model{},
		Name:       "",
		StudentNum: studentnum,
		PassWord:   password,
		Salt:       salt,
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

func JwtGenerate(id int, name string) (string, error) {
	jwtClaims := JwtClaims{
		Id:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "Token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	return token.SignedString([]byte(config.Conf.GetString("data.jwt.stSignKey")))
}

func ParseToken(token string) (JwtClaims, error) {
	jwtClaims := JwtClaims{}
	tokenStr, err := jwt.ParseWithClaims(token, &jwtClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Conf.GetString("data.jwt.stSignKey")), nil
	})
	if err == nil && !tokenStr.Valid {
		err = errors.New("invalid token")
		return jwtClaims, err
	}
	if err != nil {
		return jwtClaims, err
	}
	claims, ok := tokenStr.Claims.(*jwt.RegisteredClaims)
	if claims == nil {
		return jwtClaims, errors.New("token错误")
	}
	if !ok {
		return jwtClaims, errors.New("类型断言失败")
	}
	if claims.ExpiresAt.Before(time.Now()) {
		return jwtClaims, errors.New("token已过期")
	}
	return jwtClaims, nil
}
