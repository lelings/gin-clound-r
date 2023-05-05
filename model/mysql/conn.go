package mysql

import (
	"fmt"
	"log"

	"example.com/m/v2/lib"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		lib.SERVICE_CONFIG.Mysql_name, lib.SERVICE_CONFIG.Mysql_password,
		lib.SERVICE_CONFIG.Mysql_host, lib.SERVICE_CONFIG.Mysql_port,
		lib.SERVICE_CONFIG.Mysql_database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("[InitDB error]: open mysql error:", err)
	}
	DB = db

}
