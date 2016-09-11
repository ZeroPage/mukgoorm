package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"fmt"
)

func setPassword(adminPassword string, readOnlyPassword string){
	file1, _ := os.Create("test.txt")        // hello1.txt 파일 생성
	defer file1.Close()                        // main 함수가 끝나기 직전에 파일을 닫음
	fmt.Fprint(file1, adminPassword + " " + readOnlyPassword)
}
func authorize(password string) string{
	var adminPassword string
	var readOnlyPassword string
	file1, _ := os.Open("test.txt")    // hello2.txt 파일 열기
	defer file1.Close()                  // main 함수가 끝나기 직전에 파일을 닫음
	fmt.Fscanln(file1, &adminPassword, &readOnlyPassword)
	if strings.Compare(adminPassword, password) == 0{
		fmt.Print("admin\n")
		return "admin"
	}
	if strings.Compare(readOnlyPassword, password) == 0{
		fmt.Print("read Only\n")
		return "readOnlyPassword"
	}
	return "fail"
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "inputPassword.tmpl", gin.H{})
	})
	r.GET("/set-login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "setPassword.tmpl", gin.H{})
	})
	r.POST("/login-data", func(c *gin.Context) {
				//user := c.PostForm("user")
				password := c.PostForm("password")
				authorize(password)
  })

	r.POST("/set-password", func(c *gin.Context) {
				adminPassword := c.PostForm("adminPassword")
				readOnlyPassword := c.PostForm("readOnlyPassword")
				fmt.Print(adminPassword + " " + readOnlyPassword + "\n")
				setPassword(adminPassword, readOnlyPassword)
	})

	r.GET("/list", func(c *gin.Context) {
		files, err := ioutil.ReadDir("tmp/dat")
		if err != nil {
			panic(err)
		}

		c.HTML(http.StatusOK, "list.tmpl", gin.H{
			"files": files,
		})
	})

	r.Run()
}
