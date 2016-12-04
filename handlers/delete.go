package handlers

import (
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/zeropage/mukgoorm/setting"
)

// Pre: target exist in rootDir
// Post: target is deleted from rootDir
func Delete(c *gin.Context) {
	target := c.Query("dir")
	rootDir := setting.GetDirectory()

	if rootDir.ValidDir(target) == false {
		log.Warnf("File Not Exist: %s", target)
		c.JSON(http.StatusNotFound, gin.H{"error": "File/Path don't exist."})
	}

	err := os.Remove(target)
	if err != nil {
		log.Warn(err)
		c.JSON(http.StatusNotFound, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{})
}
