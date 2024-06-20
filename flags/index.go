package flags

import sys_flag "flag"

type Option struct {
	// version bool
	DB bool
}

// 解析命令行参数
func Parse() Option {
	// version := sys_flag.Bool("v", false, "版本")
	db := sys_flag.Bool("db", false, "初始化数据库")
	sys_flag.Parse()
	return Option{
		// version: *version,
		DB: *db,
	}
}

// 是否停止web项目
func IsWebStop(o Option) bool {
	if o.DB {
		return false
	}
	// if o.version {
	// 	return false
	// }
	return true
}

// 根据命令行执行不同方法
func SwitchOption(o Option) {
	if o.DB {
		MakeMigrations()
	}
	// if o.version {
	// 	Version()
	// }
}
