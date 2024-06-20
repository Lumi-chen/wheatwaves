package flags

import (
	"wheatwaves/global"
	"wheatwaves/models"
)

func MakeMigrations() {
	// fmt.Println("MakeMigrations")

	err := global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.UserModel{},
			&models.ArticleModel{},
			&models.ColumnModel{},
			&models.ImageModel{},
			&models.TagModel{},
			&models.TouristModel{},
		)

	if err != nil {
		global.Log.Error("[error] 数据库表结构生成失败")
		return
	}
	global.Log.Info("[Success] 数据库表结构生成成功")
	// 如果自定义了连接表，可以通过该方法修改连接表
	// global.DB.SetupJoinTable(&models.UserModel{}, "Articles", &models.User2Article{})
}
