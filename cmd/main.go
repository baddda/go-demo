package main

import (
	"tasko/internal/router"
	"tasko/internal/util"
)

func main() {
	util.ConnectDatabase()
	router := router.SetupRouter()
	router.Run("localhost:8080")
}
