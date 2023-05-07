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

//get the all files in the folder
func GetParentFolder(folderId int, storeId int) []FileFolder {
	var folder []FileFolder
	mysql.DB.Where("parent_folder_id = ? and file_store_id = ?", folderId, storeId).Find(&folder)
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
func GetCurrentAllParent(folder FileFolder, folders []FileFolder) []FileFolder {
	if folder.ParentFloderId != 0 {
		folder = GetParentFolderId(folder.ParentFloderId)
		folders = append(folders, folder)
		folders = GetCurrentAllParent(folder, folders)
	}
	for i, j := 0, len(folders)-1; i < j; i, j = i+1, j-1 {
		folders[i], folders[j] = folders[j], folders[i]
	}
	return folders
}

//get user's folder count
func GetFolderCount(storeId int) int {
	var count int64
	mysql.DB.Model(&FileFolder{}).Where("store_id = ?", storeId).Count(&count)
	return int(count)
}

//delete the folder and the files in the folder
func DeleteFolder(folderId int) {
	mysql.DB.Delete(&FileFolder{}, folderId)

	mysql.DB.Where("folder_id = ?", folderId).Delete(&MyFile{})

	var fileFolders []FileFolder
	mysql.DB.Find(&fileFolders, "parent_folder_id = ?", folderId)
	mysql.DB.Where("parent_folder_id = ?", folderId).Delete(&FileFolder{})

	for _, fileFolder := range fileFolders {
		DeleteFolder(fileFolder.Id)
	}
}

//update the folder's name
func UpdateFolderName(folderId int, folderName string) {
	mysql.DB.Model(&FileFolder{}).Where("id = ?", folderId).Update("file_floder_name", folderName)
}
