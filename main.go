package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zeropage/mukgoorm/cmd"
	"github.com/zeropage/mukgoorm/setting"
)

// When starting server directory parameter is needed. Else error occurs.
// Run Command:
//	go run main.go --dir tmp/dat
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
		c.Redirect(http.StatusSeeOther, "/login")
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
		c.Redirect(http.StatusSeeOther, "/login")
	}
}

func checkAuthority(c *gin.Context) { //check admin ,otherwise redirect to login or list
	session := sessions.Default(c)
	if session.Get("authority") == nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}
	switch Grant(session.Get("authority").(int)) {
	case FAIL:
		c.Redirect(http.StatusSeeOther, "/login")
	case ADMIN:
		session.Options(sessions.Options{MaxAge: SESSION_EXPIRE_TIME})
		return
	case READ_ONLY:
		session.Options(sessions.Options{MaxAge: SESSION_EXPIRE_TIME})
		c.Redirect(http.StatusSeeOther, "/list")
	}
}

func main() {
	cmd.RootCmd.Execute()
	r := NewEngine()
	r.Run()
}

func NewEngine() *gin.Engine {
	shareDir := setting.GetDirectory()
	sharePassword := setting.GetPassword()
	r := gin.Default()

	r.LoadHTMLGlob("templates/*/*.tmpl")

	store := sessions.NewCookieStore([]byte("secret"))

	r.Use(sessions.Sessions("mysession", store))

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "input_password.tmpl", gin.H{})
	})

	r.POST("/login", func(c *gin.Context) {
		password := c.PostForm("password")

		var authority Grant
		switch password {
		case sharePassword.AdminPassword:
			authority = ADMIN
		case sharePassword.ReadOnlyPassword:
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
		sharePassword.AdminPassword = c.PostForm("adminPassword")
		sharePassword.ReadOnlyPassword = c.PostForm("readOnlyPassword")
		c.Redirect(http.StatusMovedPermanently, "/login")
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

	return r
}
