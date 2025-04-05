package controller

import (
	"net/http"
	"tasko/internal/model"

	"github.com/gin-gonic/gin"
)

var tasks = []model.Task{
	{ID: "1", Description: "Buy bread"},
	{ID: "2", Description: "Toast bread"},
	{ID: "3", Description: "Eat bread"},
}

func GetTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

func PostTask(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": r})
		}
	}()

	var newTask model.Task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	if newTask.Description == "" {
		panic("Description is required")
	}

	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}
