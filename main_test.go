package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"

	"github.com/zeropage/mukgoorm/setting"

	"github.com/stretchr/testify/assert"
)

func initialize() {
	setting := setting.GetDirectory()
	setting.Path = "tmp/dat"
}

var session string
var once sync.Once

func initializeSession() {
	once.Do(func() {
		initialize()

		r := NewEngine()

		data := url.Values{}
		data.Set("password", "admin") // TODO password can be changed
		req, _ := http.NewRequest("POST", "/login", bytes.NewBufferString(data.Encode()))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		session = w.Header().Get("Set-Cookie")
	})
}

// This code came from gin-gonic/gin/routes_test.go
func PerformRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	initialize()

	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func PerformRequestWithSession(r http.Handler, method, path string) *httptest.ResponseRecorder {
	initializeSession()

	req, _ := http.NewRequest(method, path, nil)
	req.Header.Add("Cookie", session)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestAuthoritySuccess(t *testing.T) {
	r := NewEngine()

	w := PerformRequestWithSession(r, "GET", "/list")
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAuthorityFail(t *testing.T) {
	r := NewEngine()

	w := PerformRequest(r, "GET", "/list")
	assert.Equal(t, http.StatusSeeOther, w.Code)
}

func TestRoutes(t *testing.T) {
	r := NewEngine()

	w := PerformRequest(r, "GET", "/list")
	assert.NotEqual(t, http.StatusNotFound, w.Code)

	w = PerformRequest(r, "GET", "/down?fn=hello1.txt")
	assert.NotEqual(t, http.StatusNotFound, w.Code)

	w = PerformRequest(r, "GET", "/login")
	assert.NotEqual(t, http.StatusNotFound, w.Code)

	w = PerformRequest(r, "POST", "/login")
	assert.NotEqual(t, http.StatusNotFound, w.Code)

	w = PerformRequest(r, "GET", "/set-password")
	assert.NotEqual(t, http.StatusNotFound, w.Code)

	w = PerformRequest(r, "POST", "/set-password")
	assert.NotEqual(t, http.StatusNotFound, w.Code)
}
