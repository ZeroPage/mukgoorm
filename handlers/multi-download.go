package handlers
import (
	"net/http"
	"github.com/gin-gonic/gin"

	"os"
	"archive/zip"
	"path/filepath"
	"strings"
	"io"
	"io/ioutil"

		log "github.com/Sirupsen/logrus"
)
func MultiDownload(c *gin.Context) {
	c.Request.ParseForm()
	f := c.Request.PostForm["chk_info"]
	fN := c.Request.PostForm["fileName"]
	zipit(f, fN[0])
	fileName := fN[0] + ".zip"
	defer os.Remove(fileName)
	filedata, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Error(err)
		c.HTML(http.StatusNotFound, "errors/404.tmpl", gin.H{})
	}
	_, fileName = filepath.Split(fileName)
	c.Writer.Header().Set("content-disposition", "attachment; filename=" + fileName)
	c.Data(http.StatusOK, "application/octet-stream", filedata)
	c.Redirect(http.StatusSeeOther, "/list")
}

func zipit(source []string, target string) error {
	zipfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()
	for _, v := range source{
		_, err := os.Stat(v)
		if err != nil {
			return nil
		}
	}
	var baseDir string

	for _, v := range source {
		info, err := os.Stat(v)
		if err != nil {
			return nil
		}
		if info.IsDir() {
			baseDir = filepath.Base(v)
		}	else{
			baseDir = ""
		}
		filepath.Walk(v, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			header, err := zip.FileInfoHeader(info)
			if err != nil {
				return err
			}

			if baseDir != "" {
				header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, v))
			}

			if info.IsDir() {
				header.Name += "/"
			} else {
				header.Method = zip.Deflate
			}

			writer, err := archive.CreateHeader(header)
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
			return err
		})
	}
	return err
}
