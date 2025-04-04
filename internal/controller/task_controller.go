package controller

import (
	"net/http"
	"tasko/internal/model"

	"github.com/gin-gonic/gin"
)

var tasks = []model.Task{
	{ID: "1", Description: "Blue Train"},
	{ID: "2", Description: "Jeru"},
	{ID: "3", Description: "Sarah Vaughan and Clifford Brown"},
}

func GetTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

func PostTask(c *gin.Context) {
	var newTask model.Task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	if newTask.Description == "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Description is required"})
		return
	}

	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}
