package handlers

import (
	"net/http"
	"os"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/zeropage/mukgoorm/setting"
)

func List(c *gin.Context) {
	shareDir := setting.GetDirectory()

	sharedPath := c.Query("dir")
	if sharedPath == "" {
		sharedPath = shareDir.Path
	} else if !shareDir.ValidDir(sharedPath) {
		log.Infof("Invalid directory access: %s", sharedPath)
		c.HTML(http.StatusNotFound, "errors/404.tmpl", gin.H{})
	}

	files, err := getFileInfoAndPath(sharedPath)
	if err != nil {
		log.Error(err)
		c.HTML(http.StatusNotFound, "errors/404.tmpl", gin.H{})
	}
	c.HTML(http.StatusOK, "common/list.tmpl", gin.H{
		"files": files,
	})
}

func getFileInfoAndPath(root string) (*[]FilePathInfo, error) {
	files := []FilePathInfo{}
	err := filepath.Walk(root, filepath.WalkFunc(func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip base directory
		if path == root {
			return nil
		}

		files = append(files, FilePathInfo{info, path})
		if info.IsDir() {
			return filepath.SkipDir
		}
		return err
	}))
	return &files, err
}

type FilePathInfo struct {
	File os.FileInfo
	Path string
}
