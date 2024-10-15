package dao

import (
	"github.com/Camelia-hu/tuan/config"
	"github.com/Camelia-hu/tuan/modules"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func DB_Init() {
	dsn := config.Conf.GetString("data.mysql.dsn")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln("db open err :  ", err)
	}
	err = db.AutoMigrate(&modules.User{})
	if err != nil {
		log.Println(err)
	}
	DB = db
}
