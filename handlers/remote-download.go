package handlers

import (
	"os"
	"net/http"
	"io"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)
func RemoteDownload(c *gin.Context) {
	url := c.PostForm("url")
	tokens := strings.Split(url, "/")
	tokens = strings.Split(tokens[len(tokens) -1], "?")
	fileName := time.Now().Format("2006-01-02150405") +"_"+ tokens[0]
	err := downloadFile(fileName, url)
	if err != nil {
		panic(err)
	}
	c.Redirect(http.StatusSeeOther, "/list")
}

func downloadFile(filename string, url string) (err error) {
	out, err := os.Create("./tmp/dat/"+filename)
	if err != nil  {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil  {
		return err
	}

	return nil
}
