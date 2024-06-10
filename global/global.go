package global

import (
	"wheatwaves/config"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// 将配置文件的信息放在公共变量里
var (
	Config *config.Config
	DB     *gorm.DB
	Log    *logrus.Logger
	// 使用指针便于修改
)
