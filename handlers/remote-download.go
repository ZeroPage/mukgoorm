package handlers

import (
	"io"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zeropage/mukgoorm/image"
	"github.com/zeropage/mukgoorm/setting"
)

func RemoteDownload(c *gin.Context) {
	url := c.PostForm("url")
	tokens := strings.Split(url, "/")
	tokens = strings.Split(tokens[len(tokens)-1], "?")
	fileName := time.Now().Format("2006-01-02150405") + "_" + tokens[0]
	filePath := path.Join(setting.GetDirectory().Path, fileName)
	_, err := downloadFile(filePath, url)
	if err != nil {
		panic(err)
	}

	go image.Resize(filePath, 300)

	c.Redirect(http.StatusSeeOther, "/list")
}

func downloadFile(filePath string, url string) (string, error) {
	out, err := os.Create(filePath)
	if err != nil {
		return filePath, err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return filePath, err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return filePath, err
	}

	return filePath, nil
}
