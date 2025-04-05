package controller_test

import (
	"net/http"
	"net/http/httptest"
	"tasko/internal/controller"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetUserByID(t *testing.T) {
	router := gin.Default()
	router.GET("/api/tasks", controller.GetTasks)

	req, err := http.NewRequest("GET", "/api/tasks", nil)
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "[\n    {\n        \"id\": \"1\",\n        \"description\": \"Buy bread\"\n    },\n    {\n        \"id\": \"2\",\n        \"description\": \"Toast bread\"\n    },\n    {\n        \"id\": \"3\",\n        \"description\": \"Eat bread\"\n    }\n]", rec.Body.String())
}
