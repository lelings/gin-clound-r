package model

import (
	"time"

	"example.com/m/v2/model/mysql"
)

type FileFolder struct {
	Id             int
	FileFloderName string
	ParentFloderId int
	FileStoreId    int
	Time           string
}

func CreateFolder(folderName string, parentId, storeId int) {
	fileFolder := FileFolder{
		FileFloderName: folderName,
		ParentFloderId: parentId,
		FileStoreId:    storeId,
		Time:           time.Now().Format("2006-01-02 15:04:05"),
	}
	mysql.DB.Create(&fileFolder)
}

func DeleteFolder(folderId int) {
	mysql.DB.Delete(&FileFolder{}, folderId)
}

//get the all files in the folder
func GetParentFolder(folderId int, storeId int) []FileFolder {
	var folder []FileFolder
	mysql.DB.Where("parent_folder_id = ? and store_id = ?", folderId, storeId).Find(&folder)
	return folder
}

//get the parent folder's id
func GetParentFolderId(folderId int) FileFolder {
	var folder FileFolder
	mysql.DB.Where("id = ?", folderId).First(&folder)
	return folder
}

//get the foldler's detail
func GetFolderDetail(folderId int) FileFolder {
	var folder FileFolder
	mysql.DB.Where("id = ?", folderId).First(&folder)
	return folder
}

//get current folder's all parent folder
func GetCurrentAllParent()
