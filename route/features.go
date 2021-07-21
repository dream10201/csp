package route

import (
	"github.com/csp/v2/cmd"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-gonic/gin"
	"log"
)

var r *gin.Engine

func init() {
	r = gin.Default()
	r.POST("/apply", apply)
	staticFS := assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
		Prefix:    "frontend",
		Fallback:  "index.html",
	}
	r.StaticFS("/", &staticFS)
}
func apply(context *gin.Context) {
	user := context.PostForm("user")
	pwd := context.PostForm("pwd")
	newPwd := context.PostForm("newPwd")
	status := 0
	if cmd.CheckOldPassword(user, pwd) && cmd.ChangePassword(user, newPwd) {
		status = 1
	}
	context.JSON(200, gin.H{
		"status": status,
	})
}
func Run(addr string) {
	err := r.Run(addr)
	if err != nil {
		log.Println(err.Error())
	}
}
