package core

import (
	"fmt"
	"log"
	"time"
	"wheatwaves/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 初始化gorm连接
func InitGorm() *gorm.DB {
	if global.Config.MySql.Host == "" {
		log.Println("未配置数据库连接信息，取消gorm连接")
		return nil
	}

	dsn := global.Config.MySql.Dsn() // 获取拼接好的连接字符串
	var mysqlLogger logger.Interface
	if global.Config.System.Env == "dev" {
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		log.Fatalf(fmt.Sprintf("[%s] mysql连接失败", dsn))
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)              // 最大可容纳数
	sqlDB.SetConnMaxLifetime(time.Hour * 4) // 连接最大复用时间，不超过mysql的wait_timeout
	return db
}
