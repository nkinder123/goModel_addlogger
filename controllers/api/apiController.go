package api

import "github.com/gin-gonic/gin"

// api接口controller
type ApiController struct {
}

func (con ApiController) Api(c *gin.Context) {
	c.JSON(200, "api_api")
}
