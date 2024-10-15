package main

import (
	"github.com/Camelia-hu/tuan/config"
	"github.com/Camelia-hu/tuan/dao"
	"github.com/Camelia-hu/tuan/routers"
)

func main() {
	config.ViperInit()
	dao.DB_Init()
	routers.RoutersInit()
}
