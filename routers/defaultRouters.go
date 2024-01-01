package routers

import (
	"Gin/controllers/front"
	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/front", front.FrontController{}.Front1)
	}
}
