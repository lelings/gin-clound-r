package handler

import (
	"fmt"
	"net/http"

	"example.com/m/v2/lib"
	"example.com/m/v2/middleware"
	"example.com/m/v2/model"
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
	token := middleware.SpawnToken(user.Id, user.UserName)
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
	// username := c.PostForm("username")
	// password := c.PostForm("password")

}
