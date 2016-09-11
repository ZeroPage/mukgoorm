package main

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)
type Grant int
const (
	FAIL Grant = iota
	ADMIN
	READ_ONLY
)
const SESSION_EXPIRE_TIME int = 1800

func checkLogin(c *gin.Context) { //check login
	session := sessions.Default(c)
	if session.Get("authority") == nil {
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}
	switch Grant(session.Get("authority").(int)) {
		//Grant(session.Get("authority")) will cause error: "cannot convert session.Get("authority") (type interface {}) to type Grant: need type assertion"
		//session.Get("authority").(int) will cause error: "invalid case ADMIN(and READ_ONLY) in switch on session.Get("authority").(int) (mismatched types Grant and int)"
	case ADMIN:
		return
	case READ_ONLY:
		return
	default:
		c.Redirect(http.StatusMovedPermanently, "/login")
	}
}

func checkAuthority(c *gin.Context) { //check admin ,otherwise redirect to login or list
	session := sessions.Default(c)
	if session.Get("authority") == nil {
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}
	switch Grant(session.Get("authority").(int)) {
	case FAIL:
		c.Redirect(http.StatusMovedPermanently, "/login")
	case ADMIN:
		session.Options(sessions.Options{MaxAge: SESSION_EXPIRE_TIME})
		return
	case READ_ONLY:
		session.Options(sessions.Options{MaxAge: SESSION_EXPIRE_TIME})
		c.Redirect(http.StatusMovedPermanently, "/list")
	}
}

func main() {
	r := gin.Default()

	store := sessions.NewCookieStore([]byte("secret"))

	r.Use(sessions.Sessions("mysession", store))

	r.LoadHTMLGlob("templates/*")

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "input_password.tmpl", gin.H{})
	})

	r.POST("/login", func(c *gin.Context) {
		password := c.PostForm("password")
		var adminPassword, readOnlyPassword string
		file1, _ := os.Open("test.txt")
		fmt.Fscanln(file1, &adminPassword, &readOnlyPassword)
		file1.Close()
		var authority Grant
		switch password {
		case adminPassword:
			authority = ADMIN
		case readOnlyPassword:
			authority = READ_ONLY
		default:
			authority = FAIL
		}
		session := sessions.Default(c)
		session.Set("authority", int(authority))
		//if you just put authority which is Grant type, then session save nil....
		session.Options(sessions.Options{MaxAge: SESSION_EXPIRE_TIME})
		session.Save()
		c.Redirect(http.StatusMovedPermanently, "/list")
	})

	r.GET("/set-password", func(c *gin.Context) {
		checkAuthority(c)
		c.HTML(http.StatusOK, "set_password.tmpl", gin.H{})
	})

	r.POST("/set-password", func(c *gin.Context) {
		file1, _ := os.Create("test.txt")
		fmt.Fprint(file1, c.PostForm("adminPassword")+" "+c.PostForm("readOnlyPassword"))
		file1.Close()
		c.Redirect(http.StatusMovedPermanently, "/login")
	})

	r.GET("/list", func(c *gin.Context) {
		checkLogin(c)
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
