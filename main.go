package main

import (
	"wheatwaves/core"
	"wheatwaves/flags"
	"wheatwaves/global"
	"wheatwaves/routers"
)

func main() {
	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	// 测试：global.Log.Warn("警告啦")
	// 连接数据库
	global.DB = core.InitGorm()
	// fmt.Println("\n", global.Config)
	// fmt.Println("\n", global.DB)
	// 迁移表结构
	option := flags.Parse()
	if !flags.IsWebStop(option) {
		flags.SwitchOption(option)
		return
	}
	router := routers.InitRouter()
	router.Run(global.Config.System.Addr())
}
