package handlers

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zeropage/mukgoorm/grant"
)

func LoginForm(c *gin.Context) {
	c.HTML(http.StatusOK, "authority/input_password.tmpl", gin.H{})
}

func Login(c *gin.Context) {
	password := c.PostForm("password")

	authority := grant.FromPassword(password)
	session := sessions.Default(c)
	// INFO: if you just put authority which is Grant type, then session save nil....
	session.Set("authority", int(authority))
	session.Options(sessions.Options{MaxAge: SESSION_EXPIRE_TIME})
	session.Save()

	c.Redirect(http.StatusFound, "/")
}
