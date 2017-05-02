package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/zeropage/mukgoorm/setting"
)

func init() {
	dir := setting.GetDirectory()
	dir.Path = "../testdata"
}

func PerformRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestImage(t *testing.T) {
	r := gin.Default()
	r.GET("/img/:name", Image)

	w := PerformRequest(r, "GET", "/img/pic.jpg")
	assert.Equal(t, http.StatusOK, w.Code)
}
