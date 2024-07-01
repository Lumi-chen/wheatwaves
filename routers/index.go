package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// gin.SetMode(global.Config.System.Env) // gin.setMode() 用于设置运行环境
	router := gin.Default()
	SystemRouter(router)
	UserRouter(router)
	return router
}
