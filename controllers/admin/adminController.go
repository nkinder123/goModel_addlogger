package admin

import (
	"Gin/controllers/public"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 管理员操作
type AdminController struct {
	public.PublicStructController
}

func (con AdminController) ListA(context *gin.Context) {
	usrename, _ := context.Get("username")
	fmt.Printf("this is admin_username:%v", usrename)
}
