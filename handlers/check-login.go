package handlers

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zeropage/mukgoorm/grant"
)

func CheckLogin(c *gin.Context) {
	session := sessions.Default(c)
	auth := grant.FromSession(session.Get("authority"))

	authorized, err := grant.AuthorityExist(auth)
	if !authorized {
		c.Redirect(http.StatusSeeOther, "/login")
		c.Abort()
	}
	if err != nil {
		panic(err)
		c.Redirect(http.StatusSeeOther, "/login")
		c.Abort()
	}
}
