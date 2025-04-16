package controller

import (
	"log"
	"net/http"
	"strconv"
	"tasko/internal/model"
	"tasko/internal/util"

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
		c.AbortWithError(http.StatusInternalServerError, err)
	}
}

func getTasksFromDB() ([]model.Task, error) {
	log.Println("Fetching tasks from DB" + strconv.FormatBool((util.DBCon == nil)))
	rows, err := util.DBCon.Query("SELECT * FROM task")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return []model.Task{}, nil
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
