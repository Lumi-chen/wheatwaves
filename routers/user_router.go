package routers

import (
	"wheatwaves/api"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	userApi := api.ApiGroupApp.UserApi
	userGroup := router.Group("user")
	userGroup.POST("/signUpUser", userApi.SignUpUser)
	userGroup.POST("/signInUser", userApi.SignInUser)
}
