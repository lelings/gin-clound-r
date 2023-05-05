package model

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
	Type           string //文件类型
	PostFix        string //文件后缀
}
