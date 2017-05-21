package handlers

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func Info(c *gin.Context) {
	fileName := c.Query("dir")
	fileinfo, err := os.Stat(fileName)
	if err != nil {
		panic(err)
	}

	c.HTML(http.StatusOK, "common/info.tmpl", gin.H{
		"filename":   fileinfo.Name(),
		"directory":  strings.Split(fileName, fileinfo.Name())[0],
		"size":       fileinfo.Size(),
		"type":       fileinfo.Mode(), // FIXME : This is not type that we want.
		"overwriten": fileinfo.ModTime(),
	})
}
