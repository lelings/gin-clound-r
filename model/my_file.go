package model

import (
	"path"
	"strconv"
	"strings"
	"time"

	"example.com/m/v2/model/mysql"
	"example.com/m/v2/util"
)

type MyFile struct {
	Id             int    //ID
	FileName       string //文件名
	FileHash       string //文件哈希值
	FileStoreId    int    //文件仓库ID
	FilePath       string //文件路径
	ParentFolderId int    //父文件夹ID
	DownloadNum    int    //下载次数
	UploadTime     string //上传时间
	Size           int64  //文件大小
	SizeStr        string //文件大小单位
	Type           int    //文件类型
	PostFix        string //文件后缀
}

func CreateFile(file, fileHash string, fileStoreId int, fid string, size int64) {
	var sizeStr string
	//文件后缀
	fileSuffix := path.Ext(file)
	fileName := file[:len(file)-len(fileSuffix)]
	if size < 1048576 {
		sizeStr = strconv.FormatInt((size/1024), 10) + "KB"
	} else {
		sizeStr = strconv.FormatInt((size/102400), 10) + "MB"
	}
	fId, _ := strconv.Atoi(fid)
	myFile := MyFile{
		FileName:       fileName,
		FileHash:       fileHash,
		FileStoreId:    fileStoreId,
		FilePath:       "",
		DownloadNum:    0,
		UploadTime:     time.Now().Format("2006-01-02 15:04:05"),
		ParentFolderId: fId,
		Size:           size / 1024,
		SizeStr:        sizeStr,
		Type:           util.GetFileType(fileSuffix),
		PostFix:        strings.ToLower(fileSuffix),
	}
	mysql.DB.Create(&myFile)
}

func GetUserFile(folderId, storeId int) []MyFile {
	var myFiles = []MyFile{}
	mysql.DB.Find(&myFiles, "parent_folder_id = ? and store_id = ?", folderId, storeId)
	return myFiles
}

func GetUserFileCount(storeId int) int {
	var cnt int64
	mysql.DB.Find(&MyFile{}, "where store_id = ?", storeId).Count(&cnt)
	return int(cnt)
}

func SubSize(size int64, storeId int) {
	var fileStore FileStore
	mysql.DB.First(&fileStore, storeId)
	fileStore.CurrentSize += size / 1024
	fileStore.MaxSize -= size / 1024
	mysql.DB.Save(&fileStore)
}

func GetFileByType(fileType, storeId int) (files []MyFile) {
	mysql.DB.Find(&files, "where fileType = ? and storeId = ?", fileType, storeId)
	return
}

func IsFileExist(fileName string, fid int) bool {
	var currentFile MyFile

	fileSuffix := strings.ToLower(path.Ext(fileName))
	filePrefix := fileName[:len(fileName)-len(fileSuffix)]

	mysql.DB.First(&currentFile, "where parent_folder_id = ? and file_name = ? and postfix = ?", fid, filePrefix, fileSuffix)

	if currentFile.Size > 0 {
		return true
	}

	return false
}

func GetDetailFile(storeId int) map[string]int64 {
	fileDetail := make(map[string]int64)

	var files []MyFile

	var (
		docFile   int64
		imgFile   int64
		videoFile int64
		musicFile int64
		otherFile int64
	)

	mysql.DB.Find(&files, "where file_store_id = ? and type = ?", storeId, 1).Count(&docFile)
	mysql.DB.Find(&files, "where file_store_id = ? and type = ?", storeId, 2).Count(&imgFile)
	mysql.DB.Find(&files, "where file_store_id = ? and type = ?", storeId, 3).Count(&videoFile)
	mysql.DB.Find(&files, "where file_store_id = ? and type = ?", storeId, 4).Count(&musicFile)
	mysql.DB.Find(&files, "where file_store_id = ? and type = ?", storeId, 5).Count(&otherFile)

	fileDetail["docCount"] = docFile
	fileDetail["imgCount"] = imgFile
	fileDetail["videoCount"] = videoFile
	fileDetail["musicCount"] = musicFile
	fileDetail["otherCount"] = otherFile

	return fileDetail
}

func AddDownloadNum(id int) {
	var file MyFile
	mysql.DB.First(&file, id)
	file.DownloadNum += 1
	mysql.DB.Save(&file)
}

func DeleteFile(id, fid, storeId int) {
	mysql.DB.Delete(&MyFile{}, "where id = ? and parent_file_id = ? and store_id = ?", id, fid, storeId)
}
