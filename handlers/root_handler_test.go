// Package handlers provides HTTP request handlers.
package handlers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestPing(t *testing.T) {
	// Setup
	router := gin.New()
	router.GET("/ping", Ping)

	// Request
	req, err := http.NewRequest(http.MethodGet, "/ping", nil)
	assert.NoError(t, err)

	// Response
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Assertions
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"message": "pong"}`, rec.Body.String())
}
