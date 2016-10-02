package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// This code came from gin-gonic/gin/routes_test.go
func PerformRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestRoutes(t *testing.T) {
	r := NewEngine()

	w := PerformRequest(r, "GET", "/list")
	assert.Equal(t, w.Code, http.StatusOK)

	w = PerformRequest(r, "GET", "/down?fn=hello1.txt")
	assert.Equal(t, w.Code, http.StatusOK)
}
