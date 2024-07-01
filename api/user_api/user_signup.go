package userapi

import (
	"fmt"
	"wheatwaves/global"
	"wheatwaves/models"
	"wheatwaves/models/mtype"
	"wheatwaves/utils/pwd"
	"wheatwaves/utils/response"

	"github.com/gin-gonic/gin"
)

// 用户表
type UserCreateRequest struct {
	NickName   string     `json:"nick_name" form:"nickName" binding:"required" msg:"请输入昵称"`  // 昵称
	UserName   string     `json:"user_name" form:"userName" binding:"required" msg:"请输入用户名"` // 用户名
	Password   string     `json:"password" form:"password"  binding:"required" msg:"请输入密码"`  // 密码
	Role       mtype.Role `json:"role" form:"role"`                                          // 角色权限： 1-管理员、2-博主、3-游客
	IP         string     `json:"ip" form:"ip"`                                              // IP地址
	AvatarPath string     `json:"avatar_path" form:"avatarPath"`                             // 用户头像
}

// 用户注册接口
func (UserApi) SignUpUser(c *gin.Context) {
	var u UserCreateRequest
	if err := c.Bind(&u); err != nil {
		response.Fail(err, "报错了", c)
		return
	}
	fmt.Println(u, c)
	var userModel models.UserModel
	dbErr := global.DB.Take(&userModel, "user_name = ?", u.UserName).Error
	if dbErr != nil && dbErr.Error() != "record not found" {
		// 用户名存在
		global.Log.Error("用户名已存在，请重新输入", dbErr)
		response.Fail(dbErr, "用户名已存在", c)
		return
	}
	avatarPath := "https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
	hashPwd, _ := pwd.PasswordHash(u.Password)
	global.DB.Create(&models.UserModel{
		NickName:   u.NickName,
		UserName:   u.UserName,
		Password:   hashPwd,
		AvatarPath: avatarPath,
		Role:       mtype.OrdinaryUser, // 注册默认是普通用户
		IP:         "127.0.0.0",
	})
	// if createErr != nil {
	// 	global.Log.Error("用户名创建失败，请重试", createErr)
	// 	response.Fail(createErr, "用户名创建失败，请重试", c)
	// 	return
	// }
	response.Success("用户注册成功", "success", c)

}
