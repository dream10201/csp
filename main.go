package main

import (
	"flag"
	"fmt"
	"github.com/csp/v2/cmd"
	"github.com/gin-gonic/gin"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	args := flag.String("port", "1516", "Set the args.")
	flag.Parse()
	r := gin.New()
	t, err := loadStatic()
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(t)
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/frontend/index.html", nil)
	})
	r.POST("/apply", func(context *gin.Context) {
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
	})
	r.Run(fmt.Sprintf(":%s", *args))
}
func loadStatic() (*template.Template, error) {
	t := template.New("")
	for name, file := range Assets.Files {
		defer file.Close()
		if file.IsDir() || !strings.HasSuffix(name, ".html") {
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
