package systemsetting

import (
	"wheatwaves/utils/response"

	"github.com/gin-gonic/gin"
)

func (SettingApi) SystemView(c *gin.Context) {
	// c.JSON(200, gin.H{"msg": "success"})
	response.Success("adsad", "success", c)
}
