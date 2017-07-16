package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zeropage/mukgoorm/image"
)

func Info(c *gin.Context) {
	fileName := c.Query("dir")
	fileinfo, err := os.Stat(fileName)
	if err != nil {
		panic(err)
	}

	user, _ := c.Get("user")

	// TODO file, img, dir
	fileType, location := "", ""
	if image.IsImage(fileName) {
		fileType = "image"
		location = fmt.Sprintf("/img/%s", fileinfo.Name())
	}
	c.HTML(http.StatusOK, "common/info.tmpl", gin.H{
		"src":        location,
		"type":       fileType,
		"filename":   fileinfo.Name(),
		"directory":  strings.Split(fileName, fileinfo.Name())[0],
		"size":       fileinfo.Size(),
		"overwriten": fileinfo.ModTime(),
		"user":       user,
	})
}
