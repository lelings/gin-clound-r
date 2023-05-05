package model

import (
	"time"

	"example.com/m/v2/model/mysql"
	"example.com/m/v2/util"
)

type User struct {
	Id           int
	UUid         string `gorm:"column:uuid"`
	Password     string
	FileStoreId  int
	UserName     string `gorm:"column:username"`
	Email        string
	RegisterTime time.Time
	ImagePath    string
}

func InitUser() {
	mysql.DB.AutoMigrate(&User{})
}

func CreateUser(username, password, email, image string) User {
	var user = User{
		UUid:         util.GetUUid(),
		Password:     util.GetPassword(password),
		UserName:     username,
		Email:        email,
		RegisterTime: time.Now(),
		ImagePath:    image,
	}
	mysql.DB.Create(&user)
	var fileStore = FileStore{
		UserId:      user.Id,
		CurrentSize: 0,
		MaxSize:     1048576,
	}
	mysql.DB.Create(&fileStore)
	user.FileStoreId = fileStore.Id
	mysql.DB.Save(&user)
	return user
}

func GetUserByEmail(email string) User {
	var user = User{}
	return user
}
