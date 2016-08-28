package main

import "github.com/gin-gonic/gin"
import "net/http"
import (
	"io/ioutil"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/index", func(c *gin.Context) {
		files, _ := ioutil.ReadDir("./")
		// for _, f := range files {
		// 	c.String(http.StatusOK, "%s", f.name)
		// }
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"fileList": files,
		})
	})
	router.Run(":8080")
}
