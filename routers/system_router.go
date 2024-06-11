package routers

import (
	"wheatwaves/api"

	"github.com/gin-gonic/gin"
)

func SystemRouter(router *gin.Engine) {
	settingApi := api.ApiGroupApp.SettingApi
	systemGroup := router.Group("system")
	systemGroup.GET("/init", settingApi.SystemView)
}
