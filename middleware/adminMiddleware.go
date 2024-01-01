package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 设置中间件
func PublicMiddle(context *gin.Context) {
	fmt.Println("middle--context1")
	context.Set("username", "lizhiyao")
}
