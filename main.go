package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type task struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

var tasks = []task{
	{ID: "1", Description: "Blue Train"},
	{ID: "2", Description: "Jeru"},
	{ID: "3", Description: "Sarah Vaughan and Clifford Brown"},
}

func getTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

func postTasks(c *gin.Context) {
	var newTask task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

func main() {
	router := gin.Default()
	router.GET("/tasks", getTasks)

	router.POST("/tasks", postTasks)

	router.Run("localhost:8080")
}
