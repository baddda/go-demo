package router

import (
	"tasko/internal/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// public := router.Group("/api")
	// {
	// 	public.POST("/register", controllers.Register)
	// 	public.POST("/login", controllers.Login)
	// }

	protected := router.Group("/api")
	// protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/tasks", controller.GetTasks)
		protected.POST("/tasks", controller.PostTask)
	}
	return router
}
