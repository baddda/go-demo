package controller

import (
	"net/http"
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
		c.JSON(http.StatusOK, tasks)
	case err := <-errorsCh:
		c.AbortWithError(http.StatusInternalServerError, err)
	}
}

func getTasksFromDB() ([]model.Task, error) {
	rows, err := util.DBCon.Query("SELECT * FROM task")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task

	for rows.Next() {
		var task model.Task
		if err := rows.Scan(&task.ID, &task.Description); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func PostTask(c *gin.Context) {
	var newTask model.Task

	if err := c.BindJSON(&newTask); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if newTask.Description == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Description is required"})
		return
	}

	tasksSample = append(tasksSample, newTask)
	c.JSON(http.StatusCreated, newTask)
}
