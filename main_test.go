package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zeropage/mukgoorm/setting"

	"github.com/stretchr/testify/assert"
)

func initialize() {
	setting := setting.GetDirectory()
	setting.Path = "tmp/dat"
}

// This code came from gin-gonic/gin/routes_test.go
func PerformRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestRoutes(t *testing.T) {
	initialize()

	r := NewEngine()

	w := PerformRequest(r, "GET", "/list")
	assert.Equal(t, w.Code, http.StatusMovedPermanently)

	w = PerformRequest(r, "GET", "/down?fn=hello1.txt")
	assert.Equal(t, w.Code, http.StatusOK)
}
