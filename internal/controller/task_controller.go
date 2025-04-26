package controller

import (
	"log"
	"net/http"
	"tasko/internal/model"
	"tasko/internal/util"

	"github.com/gin-gonic/gin"
)

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
		c.JSON(http.StatusOK, tasks)
	case err := <-errorsCh:
		c.AbortWithError(http.StatusInternalServerError, err)
	}
}

func getTasksFromDB() ([]model.Task, error) {
	var tasks []model.Task
	if err := util.DBCon.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func PostTask(c *gin.Context) {
	var newTask model.Task

	if err := c.BindJSON(&newTask); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if newTask.Description == "" {
		log.Println("PostTask description is required")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Description is required"})
		return
	}

	log.Printf("task %s", newTask)

	if err := util.DBCon.Create(&newTask).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, newTask)
}
