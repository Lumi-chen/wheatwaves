package api

import (
	systemsetting "wheatwaves/api/system_setting"
	userapi "wheatwaves/api/user_api"
)

type ApiGroup struct {
	SettingApi systemsetting.SettingApi
	UserApi    userapi.UserApi
}

var ApiGroupApp = new(ApiGroup)
