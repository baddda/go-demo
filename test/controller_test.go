package controller_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"tasko/internal/controller"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetTasks(t *testing.T) {
	router := gin.Default()
	router.GET("/api/tasks", controller.GetTasks)

	req, err := http.NewRequest("GET", "/api/tasks", nil)
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "[\n    {\n        \"id\": \"1\",\n        \"description\": \"Buy bread\"\n    },\n    {\n        \"id\": \"2\",\n        \"description\": \"Toast bread\"\n    },\n    {\n        \"id\": \"3\",\n        \"description\": \"Eat bread\"\n    }\n]", rec.Body.String())
}
func TestPostTask(t *testing.T) {
	router := gin.Default()
	router.POST("/api/tasks", controller.PostTask)

	newTask := `{"description": "New task"}`

	req, err := http.NewRequest("POST", "/api/tasks", strings.NewReader(newTask))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Equal(t, "{\n    \"id\": \"\",\n    \"description\": \"New task\"\n}", rec.Body.String())
}
