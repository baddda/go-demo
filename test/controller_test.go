package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"tasko/internal/router"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTasks(t *testing.T) {
	TestConnectDatabase(t)
	router := router.SetupRouter()

	req, err := http.NewRequest("GET", "/api/tasks?isAdmin=1", nil)
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "[\n    {\n        \"id\": \"1\",\n        \"description\": \"Buy bread\"\n    },\n    {\n        \"id\": \"2\",\n        \"description\": \"Toast bread\"\n    },\n    {\n        \"id\": \"3\",\n        \"description\": \"Eat bread\"\n    }\n]", rec.Body.String())
}
func TestPostTask(t *testing.T) {
	TestConnectDatabase(t)
	router := router.SetupRouter()

	newTask := `{"description": "New task"}`

	req, err := http.NewRequest("POST", "/api/tasks?isAdmin=1", strings.NewReader(newTask))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Equal(t, "{\n    \"id\": \"\",\n    \"description\": \"New task\"\n}", rec.Body.String())
}
