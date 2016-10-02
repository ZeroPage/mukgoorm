package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/zeropage/mukgoorm/cmd"
	"github.com/zeropage/mukgoorm/setting"

	"github.com/gin-gonic/gin"
)

// When starting server directory parameter is needed. Else error occurs.
// Run Command:
//	go run main.go --dir tmp/dat
func main() {
	cmd.RootCmd.Execute()

	r := NewEngine()
	r.Run()
}

func NewEngine() *gin.Engine {
	shareDir := setting.GetDirectory()
	r := gin.Default()

	r.LoadHTMLGlob("templates/*/*.tmpl")

	r.GET("/list", func(c *gin.Context) {
		files, err := ioutil.ReadDir(shareDir.Path)
		if err != nil {
			panic(err)
		}

		c.HTML(http.StatusOK, "list.tmpl", gin.H{
			"files": files,
		})
	})

	r.GET("/down", func(c *gin.Context) {
		fileName := fmt.Sprintf("%s/%s", shareDir.Path, c.Query("fn"))
		file, err := ioutil.ReadFile(fileName)
		if err != nil {
			panic(err)
		}

		c.Data(http.StatusOK, "application/octet-stream", file)
	})

	return r
}
