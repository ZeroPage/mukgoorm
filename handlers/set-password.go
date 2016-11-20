package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zeropage/mukgoorm/setting"
)

func SetPasswordForm(c *gin.Context) {
	c.HTML(http.StatusOK, "authority/set_password.tmpl", gin.H{})
}

func SetPassword(c *gin.Context) {
	sharePassword := setting.GetPassword()

	sharePassword.AdminPassword = c.PostForm("adminPassword")
	sharePassword.ReadOnlyPassword = c.PostForm("readOnlyPassword")

	c.Redirect(http.StatusSeeOther, "/login")
}
