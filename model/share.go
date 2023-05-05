package model

type share struct {
	Id     int
	UserId int
	FileId int
	Code   string //提取码
	Hash   string
}
