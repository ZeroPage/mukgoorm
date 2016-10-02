package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	r := NewEngine()

	r.Run()
}

func NewEngine() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*/*.tmpl")

	r.GET("/list", func(c *gin.Context) {
		files, err := ioutil.ReadDir("tmp/dat")
		if err != nil {
			panic(err)
		}

		c.HTML(http.StatusOK, "list.tmpl", gin.H{
			"files": files,
		})
	})

	r.GET("/down", func(c *gin.Context) {
		fileName := fmt.Sprintf("tmp/dat/%s", c.Query("fn"))
		file, err := ioutil.ReadFile(fileName)
		if err != nil {
			panic(err)
		}

		c.Data(http.StatusOK, "application/octet-stream", file)
	})

	r.GET("/info", func(c *gin.Context) {
		fileName := fmt.Sprintf("tmp/dat/%s", c.Query("fn"))
		file, err := ioutil.ReadFile(fileName)
		if err != nil {
			panic(err)
		}

		c.HTML(http.StatusOK, "info.tmpl", gin.H{
			"files": file,
		})
		// this code is just give url(ex. localhost:8080/list?fn=hello2.txt)
	})

	r.POST("/upload", func(c *gin.Context) {

		file, header, err := c.Request.FormFile("image")
		if err != nil {
			panic(err)
		}
		filename := header.Filename
		fmt.Println(header.Filename)
		out, err := os.Create("./tmp/" + filename)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
		}
		c.Redirect(http.StatusMovedPermanently, "http://localhost:8080/list")
	})

	return r
}
