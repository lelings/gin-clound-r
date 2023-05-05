package main

import (
	"example.com/m/v2/lib"
	"example.com/m/v2/model/mysql"
	"example.com/m/v2/router"
)

func main() {
	lib.InitConfig()
	mysql.InitDB()
	lib.InitRdb()

	r := router.SetUprouter()
	r.Run("localhost:8080")
}
