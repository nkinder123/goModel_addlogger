package routers

import (
	"Gin/controllers/admin"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 设置admin路由
func RouterMiddle(ctx *gin.Context) {
	fmt.Println("这是一个中间件")
}
func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin")
	{
		adminRouters.GET("/list", RouterMiddle, admin.AdminController{}.ListA)
	}
}
