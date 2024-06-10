package routers

import (
	"wheatwaves/global"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env) // gin.setMode() 用于设置运行环境
	router := gin.Default()
	router.GET("/init", func(c *gin.Context) {
		c.String(200, "xxx")
	})
	return router
}
