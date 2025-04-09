package test

import (
	"net/http"
	"tasko/internal/router"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerRun(t *testing.T) {
	router := router.SetupRouter()

	go func() {
		err := router.Run(":8080")
		assert.NoError(t, err)
	}()

	resp, err := http.Get("http://localhost:8080/api/tasks?isAdmin=1")
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
