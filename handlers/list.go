package handlers

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/zeropage/mukgoorm/path"
	"github.com/zeropage/mukgoorm/setting"
)

func List(c *gin.Context) {
	shared := setting.GetDirectory()

	dir := c.Query("dir")
	if dir == "" {
		dir = shared.Path
	} else if !shared.Valid(dir) {
		log.Warnf("Invalid directory access: %s", dir)
		c.HTML(http.StatusNotFound, "errors/404.tmpl", gin.H{})
	}

	files, err := path.PathInfoWithDirFrom(shared.Path)
	if err != nil {
		log.Error(err)
		c.HTML(http.StatusNotFound, "errors/404.tmpl", gin.H{})
	}

	c.HTML(http.StatusOK, "common/list.tmpl", gin.H{
		"files": files,
	})
}
