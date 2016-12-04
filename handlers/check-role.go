package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zeropage/mukgoorm/grant"
)

func CheckRole(roles ...grant.Grant) func(c *gin.Context) {
	return func(c *gin.Context) {
		if auth, ok := c.Get("authority"); ok {
			for _, role := range roles {
				if role == auth {
					// pass
					c.Next()
					return
				}
			}
			//fail
			c.AbortWithStatus(http.StatusForbidden)
		} else {
			// login fail
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
		}
	}
}
