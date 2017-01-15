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
	shared := setting.GetPassword()

	shared.AdminPwd = c.PostForm("admin")
	shared.ROnlyPwd = c.PostForm("readOnly")

	c.Redirect(http.StatusSeeOther, "/login")
}
