package public

import "github.com/gin-gonic/gin"

// 公共部分接口
type PublicStructController struct {
}

func (con PublicStructController) PublicAction(context *gin.Context) {
	context.JSON(200, "public_context")
}
