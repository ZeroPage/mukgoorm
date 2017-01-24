package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zeropage/mukgoorm/grant"
	"github.com/zeropage/mukgoorm/session"
)

func LoginForm(c *gin.Context) {
	c.HTML(http.StatusOK, "authority/input_password.tmpl", gin.H{})
}

func Login(c *gin.Context) {
	password := c.PostForm("password")
	authority := grant.FromPassword(password)

	sess, _ := session.GlobalSessions.SessionStart(c.Writer, c.Request)
	defer sess.SessionRelease(c.Writer)
	// INFO: if you just put authority which is Grant type, then session save nil....
	sess.Set("authority", int(authority))

	c.Redirect(http.StatusFound, "/")
}
