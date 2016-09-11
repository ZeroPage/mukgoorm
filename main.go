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
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/list", func(c *gin.Context) {
		files, err := ioutil.ReadDir("tmp/dat")
		if err != nil {
			panic(err)
		}

		c.HTML(http.StatusOK, "list.tmpl", gin.H{
			"files": files,
		})
	})
	/*
		1. show file's information(name, size, width, height etc) if filename is clicked
		2. add new button
		3. add download event if each file's button is clicked
	*/

	r.GET("/down", func(c *gin.Context) {
		fileName := fmt.Sprintf("tmp/dat/%s", c.Query("fn"))
		file, err := ioutil.ReadFile(fileName)
		if err != nil {
			panic(err)
		}

		c.Data(http.StatusOK, "application/octet-stream", file)
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

	r.Run()
}
