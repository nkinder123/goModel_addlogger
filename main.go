package main

import (
	"Gin/routers"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"text/template"
	"time"
)

func UnixToTime(timestap int) string {
	fmt.Println(timestap)
	t := time.Unix(int64(timestap), 0)
	return t.Format("2021-09-01 09:09:09")
}
func Println(str1 string, str2 string) string {
	fmt.Println(str1, str2)
	return str1 + str2
}

func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
		"Println":    Println,
	})
	r.LoadHTMLGlob("templates/**/*")
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("session", store))

	routers.AdminRoutersInit(r)
	routers.ApiRoutersInit(r)
	routers.DefaultRoutersInit(r)

	//将get传值绑定到结构体上
	r.Run()
}
