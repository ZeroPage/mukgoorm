package handlers

import (
    "os"
    "net/http"
    "io"
		"github.com/gin-gonic/gin"
		"time"
)
func RemoteDownload(c *gin.Context) {
	url := c.PostForm("url")
	err := downloadFile(time.Now().Format("2006-01-02150405"), url)
	if err != nil {
		panic(err)
	}
	c.Redirect(http.StatusMovedPermanently, "http://localhost:8080/list")
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
