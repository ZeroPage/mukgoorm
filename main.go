package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main() {
	r := NewEngine()

	r.Run()
}

func NewEngine() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*/*.tmpl")
	//r.LoadHTMLGlob("templates/layout/**.tmpl")

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

	return r
}
