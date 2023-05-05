package test

import (
	"fmt"
	"testing"

	"example.com/m/v2/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestConn(t *testing.T) {
	dsn := "root:root@tcp(localhost:3306)/clound?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("open mysql error")
	}
	db.AutoMigrate(&model.User{})
	var user = model.User{}
	db.First(&user)
	fmt.Println(user)
}
