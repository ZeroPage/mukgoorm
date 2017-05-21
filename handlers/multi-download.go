package handlers
import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
)
func MultiDownload(c *gin.Context) {
	c.Request.ParseForm()
	f := c.Request.PostForm["chk_info"]
	fmt.Println(f[0])
	fmt.Println(f[1])
	fmt.Println(f[2])
	fN := c.Request.PostForm["fileName"]
	fmt.Println(fN[0]);


	c.Redirect(http.StatusSeeOther, "/list")
}
