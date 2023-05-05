package main

import "example.com/m/v2/router"

func main() {
	r := router.SetUprouter()
	r.Run("localhost:8080")
}
