package router

import (
	"example.com/m/v2/handler"

	"github.com/gin-gonic/gin"
)

func SetUprouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", handler.HealthChecker)
	return router
}
