package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zeropage/mukgoorm/path"
	"github.com/zeropage/mukgoorm/setting"
)

func Search(c *gin.Context) {
	query := c.Query("search")
	// TODO query check

	files := search(query)
	c.HTML(http.StatusOK, "search/list.tmpl", gin.H{
		"files": files,
	})
}

func search(query string) (res []path.FilePathInfo) {
	if query == "" {
		return
	}

	files, _ := path.PathInfoFrom(setting.GetDirectory().Path)

	for _, file := range *files {
		if strings.Contains(file.File.Name(), query) {
			res = append(res, file)
		}
	}

	return res
}
