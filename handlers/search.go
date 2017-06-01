package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zeropage/mukgoorm/path"
	"github.com/zeropage/mukgoorm/setting"
)

func Search(c *gin.Context) {
	// TODO query check
	query := c.Query("q")

	files := search(query)
	user, _ := c.Get("user")
	c.HTML(http.StatusOK, "common/list.tmpl", gin.H{
		"files": files,
		"user":  user,
	})
}

func search(query string) (res []path.FilePathInfo) {
	if query == "" {
		return
	}

	files, _ := path.PathInfoWithDirFrom(setting.GetDirectory().Path)
	for _, file := range *files {
		if strings.Contains(file.File.Name(), query) {
			res = append(res, file)
		}
	}

	return
}
