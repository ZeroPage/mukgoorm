package handlers

import (
	"io/ioutil"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/zeropage/mukgoorm/image"
)

func Image(c *gin.Context) {
	dir := image.ImagePath()
	fileName := c.Param("name")
	filedata, err := ioutil.ReadFile(path.Join(dir, fileName))
	if err != nil {
		panic(err)
	}

	c.Writer.Header().Set("content-disposition", "attachment; filename="+fileName)
	// TODO content type: https://en.wikipedia.org/wiki/Media_type
	c.Data(http.StatusOK, "image/png; image/jpeg", filedata)
}
