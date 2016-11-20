package handlers

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func Down(c *gin.Context) {

	fileName := c.Query("dir")
	file, err := os.Open(fileName)
	defer file.Close()

	fileinfo, err := file.Stat()
	if fileinfo.IsDir() {
		fileName, err = makeZip(fileName)
		if err != nil {
			panic(err)
		}
		defer os.Remove(fileName)
	}

	filedata, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Error(err)
		c.HTML(http.StatusNotFound, "errors/404.tmpl", gin.H{})
	}

	c.Data(http.StatusOK, "application/octet-stream", filedata)

}

func makeZip(foldername string) (string, error) {
	newfile, err := os.Create(foldername + ".zip")
	if err != nil {
		return "", err
	}
	defer newfile.Close()

	zipit := zip.NewWriter(newfile)
	defer zipit.Close()

	filepath.Walk(foldername, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.Name() == filepath.Base(foldername) {
			return nil
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = strings.TrimPrefix(path, foldername)

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}
		writer, err := zipit.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		zipfile, err := os.Open(path)
		defer zipfile.Close()
		if err != nil {
			return err
		}

		_, err = io.Copy(writer, zipfile)
		return err

	})
	return foldername + ".zip", err
}
