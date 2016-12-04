package handlers

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zeropage/mukgoorm/grant"
)

const SESSION_EXPIRE_TIME int = 1800

func CheckLogin(c *gin.Context) {
	session := sessions.Default(c)

	if auth, ok := grant.FromSession(session.Get("authority")); ok {
		session.Options(sessions.Options{MaxAge: SESSION_EXPIRE_TIME})
		c.Set("authority", auth)
		return
	} else {
		c.Redirect(http.StatusSeeOther, "/login")
		c.Abort()
	}
}
