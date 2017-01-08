package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zeropage/mukgoorm/cmd"
	"github.com/zeropage/mukgoorm/setting"
)

func initialize() {
	dir := setting.GetDirectory()
	dir.Path = "tmp/dat"

	pwd := setting.GetPassword()
	pwd.AdminPassword = "admin"
	pwd.ReadOnlyPassword = "readonly"
}

var session string
var once sync.Once

func initializeSession() {
	once.Do(func() {
		initialize()

		r := NewEngine()

		data := url.Values{}
		pwd := setting.GetPassword()
		data.Set("password", pwd.AdminPassword)
		req, _ := http.NewRequest("POST", "/login", bytes.NewBufferString(data.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

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

	w := PerformRequestWithSession(r, "GET", "/")
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAuthorityFail(t *testing.T) {
	r := NewEngine()
	w := PerformRequest(r, "GET", "/")
	assert.Equal(t, http.StatusSeeOther, w.Code)
}

func TestAllRoutesExist(t *testing.T) {
	routetests := []struct {
		method           string
		location         string
		expectStatusCode uint32
	}{
		{"GET", "/", http.StatusNotFound},
		// TODO generate from setting
		{"GET", "/down?dir=tmp/dat/hello1.txt", http.StatusNotFound},
		{"GET", "/login", http.StatusNotFound},
		{"POST", "/login", http.StatusNotFound},
		{"GET", "/set-password", http.StatusNotFound},
		{"POST", "/set-password", http.StatusNotFound},
		{"DELETE", "/delete", http.StatusNotFound},
	}

	r := NewEngine()

	for _, rt := range routetests {
		w := PerformRequest(r, rt.method, rt.location)
		assert.NotEqual(t, rt.expectStatusCode, w.Code)
	}
}

func TestListSuccess(t *testing.T) {
	r := NewEngine()
	w := PerformRequestWithSession(r, "GET", "/?dir=tmp/dat")
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestListFail(t *testing.T) {
	r := NewEngine()
	w := PerformRequestWithSession(r, "GET", "/?dir=tmp")
	assert.Equal(t, http.StatusNotFound, w.Code)

	w = PerformRequestWithSession(r, "GET", "/?dir=/")
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCheckStartOptionsFail(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotNil(t, r)
	}()

	for _, args := range [][]string{
		{"-A 1234", "-R 1234"},
		{"-A 1234", "-D tmp/dat"},
		{"-R 1234", "-D tmp/dat"}} {
		cmd.RootCmd.SetArgs(args)
		main()
	}
}
