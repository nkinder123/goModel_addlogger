package main

import (
	"Gin/logger"
	"Gin/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
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

	// init logger
	logger.LoggerInit()
	//创造路由
	r := gin.Default()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	//配置templates的方法
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
		"Println":    Println,
	})
	//加载静态页面路径
	r.LoadHTMLGlob("templates/**/*")

	//定义后端路由
	routers.AdminRoutersInit(r)
	routers.ApiRoutersInit(r)
	routers.DefaultRoutersInit(r)

	r.GET("/hello", func(c *gin.Context) {
		// 假设你有一些数据需要记录到日志中
		var (
			name = "q1mi"
			age  = 18
		)
		// 记录日志并使用zap.Xxx(key, val)记录相关字段
		zap.L().Debug("this is hello func", zap.String("user", name), zap.Int("age", age))

		c.String(http.StatusOK, "hello liwenzhou.com")
	})

	r.Run()
}
