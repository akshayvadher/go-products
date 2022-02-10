package config

import (
	"github.com/spf13/viper"
	"log"
	"strings"
)

var AppConf *AppConfig

var filePath = "config.yml"

type DbConfig struct {
	Username string
	Password string
	Host     string
	Name     string
}

type ServerConfig struct {
	Port int
}

type AppConfig struct {
	Server ServerConfig
	Db     DbConfig
}

func ReadConfig() {
	v := viper.New()
	v.SetConfigFile(filePath)
	v.AutomaticEnv()                                   // so that db.username can be overridden using DB_USERNAME env variable
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // db.username -> DB_USERNAME
	err := v.ReadInConfig()
	if err != nil {
		return
	}
	err = v.Unmarshal(&AppConf)
	if err != nil {
		log.Println("Error reading config " + err.Error())
		return
	}
}
