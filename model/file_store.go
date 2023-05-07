package model

import "example.com/m/v2/model/mysql"

type FileStore struct {
	Id          int
	UserId      int
	CurrentSize int64
	MaxSize     int64
}

//create a file store
func CreateFileStore(userId int, maxSize int64) {
	fileStore := FileStore{
		UserId:      userId,
		CurrentSize: 0,
		MaxSize:     maxSize,
	}
	mysql.DB.Create(&fileStore)
}

//get the file store by user id
func GetFileStore(userId int) FileStore {
	var fileStore FileStore
	mysql.DB.Where("user_id = ?", userId).First(&fileStore)
	return fileStore
}

//is the file store's size is enough
func IsEnoughSize(FileStoreId int, size int64) bool {
	var fileStore FileStore
	mysql.DB.Where("id = ?", FileStoreId).First(&fileStore)
	if fileStore.CurrentSize+(size/1024) > fileStore.MaxSize {
		return false
	}
	return true
}
