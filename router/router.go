package router

import (
	controller "example.com/m/v2/controller"
	"github.com/gin-gonic/gin"
)

func SetUprouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", controller.HealthChecker)
	user := router.Group("/user")
	{
		user.POST("/register/send", controller.SendEmailRegister)
		user.POST("/register", controller.Register)
		user.POST("/login", controller.Login)
	}
	return router
}
