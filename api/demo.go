package api

import (
	"project/conf"
	"project/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"

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
	re := service.WalletServiceAdmin{}
	err := c.ShouldBind(&re)
	if err != nil {
		c.JSON(200, util_eth.FAILURE)
		return
	}
	tok := fmt.Sprintf("token:login:manager:%s:*", re.Token)
	fmt.Println(tok)
	r := conf.RedisClient.Keys(tok)
	fmt.Println(r.Val())
	if len(r.Val()) == 0 {
		c.JSON(200, util_eth.NO_ACCESS)
		util_eth.Logzap.Error("fial..3")
		return  //todo u must unremark it
	}

	ree := re.AddAdmin()
	c.JSON(200, ree)

}
