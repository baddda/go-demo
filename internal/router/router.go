package router

import (
	"tasko/internal/controller"
	"tasko/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.ErrorHandlerMiddleware())
	router.Use(middleware.LoggingMiddleware())
	//router.Use(gin.Recovery())

	// public := router.Group("/api")
	// public.POST("/register", controllers.Register)
	// public.POST("/login", controllers.Login)

	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/tasks", controller.GetTasks)
	protected.POST("/tasks", controller.PostTask)

	return router
}
