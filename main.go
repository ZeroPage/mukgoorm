package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/GeertJohan/go.rice"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/zeropage/mukgoorm/cmd"
	"github.com/zeropage/mukgoorm/grant"
	"github.com/zeropage/mukgoorm/handlers"
	"github.com/zeropage/mukgoorm/image"
	"github.com/zeropage/mukgoorm/path"
	"github.com/zeropage/mukgoorm/setting"
)

// When starting server directory parameter is needed. Else error occurs.
// Run Command:
//	go run main.go -D tmp/dat -A *PASSWORD* -R *PASSWORD*
func main() {
	CheckStartOptions()
	go resizeImages()
	r := NewEngine()

	// FIXME recieve hostname or bind address
	r.Run("localhost:8080")
}

func CheckStartOptions() {
	cmd.RootCmd.Execute()
	if setting.GetPassword().AdminPwd == "" {
		log.Panic("Admin password must required")
	}
	if setting.GetPassword().ROnlyPwd == "" {
		log.Panic("ReadOnly password must required")
	}
	if dir := setting.GetDirectory(); dir.Path == "" || dir.Path == "." {
		log.Panicf("You need to set directory: %s", dir.Path)
	}
}

func NewEngine() *gin.Engine {
	r := gin.Default()
	r.SetHTMLTemplate(templates())

	r.StaticFS("/static", rice.MustFindBox("static").HTTPBox())

	r.GET("/login", handlers.LoginForm)
	r.POST("/login", handlers.Login)

	loginedRoute := r.Group("/", handlers.CheckLogin)

	loginedRoute.GET("/", handlers.CheckRole(grant.ADMIN, grant.READ_ONLY), handlers.List)

	loginedRoute.GET("/set-password", handlers.CheckRole(grant.ADMIN), handlers.SetPasswordForm)
	loginedRoute.POST("/set-password", handlers.CheckRole(grant.ADMIN), handlers.SetPassword)

	loginedRoute.GET("/list", handlers.CheckRole(grant.ADMIN, grant.READ_ONLY), handlers.List)
	loginedRoute.GET("/down", handlers.CheckRole(grant.ADMIN, grant.READ_ONLY), handlers.Down)
	loginedRoute.POST("/multi-download", handlers.CheckRole(grant.ADMIN, grant.READ_ONLY), handlers.MultiDownload)
	loginedRoute.GET("/info", handlers.CheckRole(grant.ADMIN, grant.READ_ONLY), handlers.Info)
	loginedRoute.POST("/upload", handlers.CheckRole(grant.ADMIN), handlers.Upload)
	loginedRoute.DELETE("/delete", handlers.CheckRole(grant.ADMIN), handlers.Delete)
	loginedRoute.GET("/search", handlers.CheckRole(grant.ADMIN, grant.READ_ONLY), handlers.Search)
	loginedRoute.POST("/remote-download", handlers.CheckRole(grant.ADMIN), handlers.RemoteDownload)

	r.GET("/img/:name", handlers.Image)

	return r
}

func templates() *template.Template {
	all := template.New("__main__").Funcs(template.FuncMap{})
	templateBox := rice.MustFindBox("templates")
	templateBox.Walk("/", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if info.Name()[0] == '.' {
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

func resizeImages() {
	image.MakeImageDir()

	root := setting.GetDirectory().Path
	files, _ := path.PathInfoWithDirFrom(root)
	for _, f := range *files {
		if f.File.IsDir() {
			continue
		}
		if image.IsImage(f.File.Name()) {
			image.Resize(300, f.Path)
		}
	}
}
