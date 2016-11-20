package handlers

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zeropage/mukgoorm/grant"
)

const SESSION_EXPIRE_TIME int = 1800

func CheckAuthority(c *gin.Context) {
	CheckLogin(c)

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
