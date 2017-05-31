package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zeropage/mukgoorm/session"
)

func Logout(c *gin.Context) {
	sess, _ := session.GlobalSessions.SessionStart(c.Writer, c.Request)
	defer sess.SessionRelease(c.Writer)

	sess.Delete("authority")
	println(sess.Get("authority"))

	c.JSON(http.StatusOK, gin.H{})
}
