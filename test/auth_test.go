package test

import (
	"net/http"
	"net/http/httptest"
	"tasko/internal/router"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthGranted(t *testing.T) {
	router := router.SetupRouter()

	req, err := http.NewRequest("GET", "/api/tasks", nil)
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestAuthNotGranted(t *testing.T) {
	router := router.SetupRouter()

	req, err := http.NewRequest("GET", "/api/tasks?isAdmin=1", nil)
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}
