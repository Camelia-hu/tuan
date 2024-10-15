package config

import (
	"flag"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

var Conf *viper.Viper

func ViperGet(path string) *viper.Viper {
	conf := viper.New()
	conf.SetConfigFile(path)
	err := conf.ReadInConfig()
	if err != nil {
		log.Println("viper read err : ", err)
	}
	conf.WatchConfig()
	conf.OnConfigChange(func(in fsnotify.Event) {
		log.Println("配置已更新哈")
		err = conf.ReadInConfig()
		if err != nil {
			log.Println("viper read err : ", err)
		}
	})
	return conf
}

func ViperInit() {
	env_config := flag.String("config", "config/app.yml", "")
	flag.Parse()
	conf := ViperGet(*env_config)
	Conf = conf
}
