package handlers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/zeropage/mukgoorm/image"
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

func before() {
	image.MakeImageDir()

	fileName := "pic.jpg"
	image.Resize(path.Join(setting.GetDirectory().Path, fileName), 300)
}

func after() {
	imageDir := image.ImagePath()
	if f, _ := os.Stat(imageDir); f != nil {
		os.RemoveAll(imageDir)
	}
}

func TestImage(t *testing.T) {
	before()
	defer after()

	r := gin.Default()
	r.GET("/img/:name", Image)

	w := PerformRequest(r, "GET", "/img/pic.jpg")
	assert.Equal(t, http.StatusOK, w.Code)
}
