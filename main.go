package main

import (
	"fmt"
	"html/template"
	"os"

	"github.com/GeertJohan/go.rice"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zeropage/mukgoorm/cmd"
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

	r.GET("/set-password", handlers.CheckAuthority, handlers.SetPasswordForm)
	r.POST("/set-password", handlers.CheckAuthority, handlers.SetPassword)

	r.GET("/list", handlers.CheckLogin, handlers.List)
	r.GET("/down", handlers.CheckAuthority, handlers.Down)
	r.GET("/info", handlers.Info)
	r.POST("/upload", handlers.Upload)
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
		template.Must(all.New(path).Parse(templateBox.MustString(path)))
		return nil
	})

	str := fmt.Sprintf("Loaded HTML Templates (%d):\n", len(all.Templates()))
	for _, v := range all.Templates() {
		str += fmt.Sprintf("\t - %s\n", v.Name())
	}
	fmt.Println(str)
	return all
}
