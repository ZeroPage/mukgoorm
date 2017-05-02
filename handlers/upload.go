package handlers

import (
	"io"
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/zeropage/mukgoorm/image"
)

func Upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		panic(err)
	}
	filename := header.Filename

	out, err := os.Create("./tmp/dat/" + filename)
	defer out.Close()
	if err != nil {
		log.Error(err)
	}

	_, err = io.Copy(out, file)
	if err != nil {
		log.Error(err)
	}

	if image.IsImage(out.Name()) {
		go image.Resize(300, out.Name())
	}

	c.Redirect(http.StatusSeeOther, "/list")
}
