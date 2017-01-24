package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zeropage/mukgoorm/grant"
	"github.com/zeropage/mukgoorm/session"
)

func CheckLogin(c *gin.Context) {
	sess, _ := session.GlobalSessions.SessionStart(c.Writer, c.Request)
	defer sess.SessionRelease(c.Writer)

	if auth, ok := grant.FromSession(sess.Get("authority")); ok {
		c.Set("authority", auth)
	} else {
		session.GlobalSessions.SessionDestroy(c.Writer, c.Request)
		c.Redirect(http.StatusSeeOther, "/login")
		c.Abort()
	}
}
