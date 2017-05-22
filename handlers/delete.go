package handlers

import (
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/zeropage/mukgoorm/setting"
)

// Pre: target exist in shared directory
// Post: target is deleted from shared directory
//
func Delete(c *gin.Context) {
	target := c.Query("dir")
	shared := setting.GetDirectory()

	if shared.Valid(target) == false {
		log.Warnf("File Not Exist: %s", target)
		c.JSON(http.StatusNotFound, gin.H{"error": "File/Path don't exist."})
	}

	err := os.Remove(target)
	if err != nil {
		log.Warn(err)
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
