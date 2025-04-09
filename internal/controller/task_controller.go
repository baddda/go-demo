package controller

import (
	"net/http"
	"tasko/internal/model"

	"github.com/gin-gonic/gin"
)

var tasksSample = []model.Task{
	{ID: "1", Description: "Buy bread"},
	{ID: "2", Description: "Toast bread"},
	{ID: "3", Description: "Eat bread"},
}

func GetTasks(c *gin.Context) {
	taskCh := make(chan []model.Task)
	errorsCh := make(chan error)

	go func() {
		tasks, err := getTasksFromDB()
		if err != nil {
			errorsCh <- err
			return
		}
		taskCh <- tasks
	}()

	select {
	case tasks := <-taskCh:
		c.IndentedJSON(http.StatusOK, tasks)
	case err := <-errorsCh:
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func getTasksFromDB() ([]model.Task, error) {
	return tasksSample, nil
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

	tasksSample = append(tasksSample, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}
