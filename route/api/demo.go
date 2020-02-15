package api

import (
	"demoproject/service"
	"demoproject/util"
	"github.com/gin-gonic/gin"
)

// @Summary swagger's demo
// @Description
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param   addr   body    string     true "common"
// @Param token body string true "token"
// @Success 200 {object} util.Errno
// @Failure 400 {object} util.Errno
// @Router /demo_api [post]
func DemoApi(c *gin.Context) {
	re := service.Demo{}
	err := c.ShouldBind(&re)
	if err != nil {
		c.JSON(200, util.FAILURE)
		return
	}

	ree := re.DemoService()
	c.JSON(200, ree)

}
