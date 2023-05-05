package lib

import (
	"log"

	"github.com/spf13/viper"
)

type ServiceConfig struct {
	Mysql_name     string
	Mysql_password string
	Mysql_database string
	Mysql_host     string
	Mysql_port     string
	Redis_host     string
	Redis_port     string
	Email_name     string
	Email_password string
}

var SERVICE_CONFIG ServiceConfig

func InitConfig() {
	viper.SetConfigName("config.json")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("[InitConfig error] ReadInConfig error:", err)
	}
	err = viper.Unmarshal(&SERVICE_CONFIG)
	if err != nil {
		log.Fatal("[InitConfig error] Unmarshal error:", err)
	}
}
