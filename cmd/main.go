package main

import (
	"tasko/internal/router"
)

func main() {
	router := router.SetupRouter()
	router.Run("localhost:8080")
}
