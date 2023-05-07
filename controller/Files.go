package handler

import (
	"strconv"

	"example.com/m/v2/middleware"
	"example.com/m/v2/model"
	"github.com/gin-gonic/gin"
)

// get all the files
func Files(c *gin.Context) {
	claims, _ := c.Get("claims")
	folderId1 := c.DefaultQuery("folderId", "0")
	claim := claims.(*middleware.Claims)
	user := model.GetUserByEmail(claim.Email)
	folderId, _ := strconv.Atoi(folderId1)

	//get the files in the folder
	files := model.GetUserFile(folderId, user.FileStoreId)

	//get the folders in the folder
	folders := model.GetParentFolder(folderId, user.FileStoreId)

	//get the parent folder
	folder := model.GetParentFolderId(folderId)

	//get all the parent folder at the current folder
	pFolders := model.GetCurrentAllParent(folder, []model.FileFolder{})

	//get the detail of the current folder
	folderDetail := model.GetFolderDetail(folderId)

	//get the detail of the current file store
	fileDetail := model.GetDetailFile(user.FileStoreId)

	c.JSON(200, gin.H{
		"code":         200,
		"fid":          folderId,
		"folderName":   folderDetail.FileFloderName,
		"files":        files,
		"folders":      folders,
		"folder":       folder,
		"pFolders":     pFolders,
		"folderDetail": folderDetail,
		"fileDetail":   fileDetail,
	})
}
