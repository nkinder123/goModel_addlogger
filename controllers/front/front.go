package front

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 前台用户default
type FrontController struct {
}

func (con FrontController) Front1(context *gin.Context) {
	fmt.Printf("这是一个default")
}
