package main

import (
	"fmt"
	"wheatwaves/core"
	"wheatwaves/global"
)

func main() {
	// 读取配置文件
	core.InitConf()
	global.DB = core.InitGorm()
	fmt.Println("\n", global.Config)
	fmt.Println("\n", global.DB)
}
