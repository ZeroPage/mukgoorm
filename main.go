package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"io"
	"os"
	"net/http"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/zeropage/mukgoorm/cmd"
	"github.com/zeropage/mukgoorm/grant"
	"github.com/zeropage/mukgoorm/setting"
)

const SESSION_EXPIRE_TIME int = 1800

func checkLogin(c *gin.Context) {
	session := sessions.Default(c)
	auth := grant.FromSession(session.Get("authority"))

	authorized, err := grant.AuthorityExist(auth)
	if !authorized {
		c.Redirect(http.StatusSeeOther, "/login")
	}
	if err != nil {
		panic(err)
		c.Redirect(http.StatusSeeOther, "/login")
	}
}

func checkAuthority(c *gin.Context) {
	checkLogin(c)

	session := sessions.Default(c)
	auth := grant.FromSession(session.Get("authority"))

	switch auth {
	case grant.ADMIN:
		session.Options(sessions.Options{MaxAge: SESSION_EXPIRE_TIME})
	case grant.READ_ONLY:
		session.Options(sessions.Options{MaxAge: SESSION_EXPIRE_TIME})
		c.Redirect(http.StatusSeeOther, "/list")
	}
}

// When starting server directory parameter is needed. Else error occurs.
// Run Command:
//	go run main.go --dir tmp/dat
func main() {
	cmd.RootCmd.Execute()
	r := NewEngine()
	// FIXME recieve hostname or bind address

	r.Run("localhost:8080")
}

func NewEngine() *gin.Engine {
	r := gin.Default()

	r.Static("/list", "./templates/javascript")
	r.LoadHTMLGlob("templates/*/*.tmpl")

	shareDir := setting.GetDirectory()
	sharePassword := setting.GetPassword()

	store := sessions.NewCookieStore([]byte("secret"))
	r.Use(sessions.Sessions("_sess", store))

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "input_password.tmpl", gin.H{})
	})

	r.POST("/login", func(c *gin.Context) {
		password := c.PostForm("password")

		authority := grant.FromPassword(password)
		session := sessions.Default(c)
		// INFO: if you just put authority which is Grant type, then session save nil....
		session.Set("authority", int(authority))
		session.Options(sessions.Options{MaxAge: SESSION_EXPIRE_TIME})
		session.Save()

		c.Redirect(http.StatusOK, "/list")
	})

	r.GET("/set-password", func(c *gin.Context) {
		checkAuthority(c)

		c.HTML(http.StatusOK, "set_password.tmpl", gin.H{})
	})

	r.POST("/set-password", func(c *gin.Context) {
		sharePassword.AdminPassword = c.PostForm("adminPassword")
		sharePassword.ReadOnlyPassword = c.PostForm("readOnlyPassword")

		c.Redirect(http.StatusOK, "/login")
	})

	r.GET("/list", func(c *gin.Context) {
		checkLogin(c)

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

	r.GET("/info", func(c *gin.Context) {
		fileName := fmt.Sprintf("tmp/dat/%s", c.Query("fn"))
		file, err := ioutil.ReadFile(fileName)
		if err != nil {
			panic(err)
		}

		c.HTML(http.StatusOK, "common/info.tmpl", gin.H{
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
		out, err := os.Create("./tmp/dat/" + filename)
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
