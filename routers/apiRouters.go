package routers

import (
	"Gin/controllers/api"
	"github.com/gin-gonic/gin"
)

// api路由
func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/api", api.ApiController{}.Api)
	}
}
