package handler

import (
	"fmt"
	"net/http"

	"example.com/m/v2/lib"
	"example.com/m/v2/middleware"
	"example.com/m/v2/model"
	"example.com/m/v2/util"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")
	code := c.PostForm("code")
	if model.GetUserByEmail(email).Id != 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "用户已存在",
		})
		return
	}

	if lib.ParseCode(code) == false {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "验证码错误",
		})
		return
	}

	user := model.CreateUser(username, password, email, "")
	token := middleware.SpawnToken(user.Id, user.UserName, user.Email)
	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"msg":   "创建成功",
		"token": token,
	})
}

func SendEmailRegister(c *gin.Context) {
	email := c.PostForm("email")
	fmt.Println(email)
	lib.SendEmail(email)
	c.JSON(http.StatusOK, gin.H{
		"msg": "sended",
	})
}

func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	user := model.GetUserByEmail(email)
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "用户不存在",
		})
		return
	}
	if user.Password != util.GetPassword(password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 401,
			"msg":  "密码错误",
		})
		return
	}

	token := middleware.SpawnToken(user.Id, user.UserName, user.Email)
	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"msg":   "登录成功",
		"token": token,
	})
}
