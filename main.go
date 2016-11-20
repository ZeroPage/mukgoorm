package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/GeertJohan/go.rice"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zeropage/mukgoorm/cmd"
	"github.com/zeropage/mukgoorm/grant"
	"github.com/zeropage/mukgoorm/handlers"
	"github.com/zeropage/mukgoorm/setting"
)

// When starting server directory parameter is needed. Else error occurs.
// Run Command:
//	go run main.go -D tmp/dat -A *PASSWORD* -R *PASSWORD*
func main() {
	cmd.RootCmd.Execute()
	if setting.GetPassword().AdminPassword == "" {
		log.Fatal("Admin password must required")
	}
	if setting.GetPassword().ReadOnlyPassword == "" {
		log.Fatal("Admin password must required")
	}
	r := NewEngine()
	// FIXME recieve hostname or bind address

	r.Run("localhost:8080")
}

func NewEngine() *gin.Engine {
	r := gin.Default()
	r.SetHTMLTemplate(templates())

	r.StaticFS("/static", rice.MustFindBox("static").HTTPBox())

	store := sessions.NewCookieStore([]byte("secret"))
	r.Use(sessions.Sessions("_sess", store))

	r.GET("/login", handlers.LoginForm)
	r.POST("/login", handlers.Login)

	loginedRoute := r.Group("/", handlers.CheckLogin)

	loginedRoute.GET("/set-password", handlers.CheckRole(grant.ADMIN), handlers.SetPasswordForm)
	loginedRoute.POST("/set-password", handlers.CheckRole(grant.ADMIN), handlers.SetPassword)

	loginedRoute.GET("/list", handlers.CheckRole(grant.ADMIN, grant.READ_ONLY), handlers.List)
	loginedRoute.GET("/down", handlers.CheckRole(grant.ADMIN, grant.READ_ONLY), handlers.Down)
	loginedRoute.GET("/info", handlers.CheckRole(grant.ADMIN, grant.READ_ONLY), handlers.Info)
	loginedRoute.POST("/upload", handlers.CheckRole(grant.ADMIN), handlers.Upload)

	return r
}
func templates() *template.Template {
	all := template.New("__main__").Funcs(template.FuncMap{})
	templateBox := rice.MustFindBox("templates")
	templateBox.Walk("/", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if path[0] == '.' {
			return nil
		}
		slashedPath := filepath.ToSlash(path)
		template.Must(all.New(slashedPath).Parse(templateBox.MustString(path)))
		return nil
	})

	str := fmt.Sprintf("Loaded HTML Templates (%d):\n", len(all.Templates()))
	for _, v := range all.Templates() {
		str += fmt.Sprintf("\t - %s\n", v.Name())
	}
	fmt.Println(str)
	return all
}
