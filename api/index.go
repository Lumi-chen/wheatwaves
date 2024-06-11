package api

import systemsetting "wheatwaves/api/system_setting"

type ApiGroup struct {
	SettingApi systemsetting.SettingApi
}

var ApiGroupApp = new(ApiGroup)
