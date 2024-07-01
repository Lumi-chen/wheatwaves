package userapi

import (
	"fmt"
	"wheatwaves/global"
	"wheatwaves/models"
	"wheatwaves/utils/jwts"
	"wheatwaves/utils/pwd"
	"wheatwaves/utils/response"

	"github.com/gin-gonic/gin"
)

type UserSignInRequest struct {
	UserName string `json:"user_name" form:"userName" binding:"required" msg:"请输入用户名"`
	PassWord string `json:"password" form:"password" binding:"required" msg:"请输入密码"`
}

func (UserApi) SignInUser(c *gin.Context) {
	var u UserSignInRequest
	err := c.Bind(&u)
	if err != nil {
		response.Fail(err, "用户名密码错误！请重新输入。", c)
		return
	}
	var userModel models.UserModel
	dberr := global.DB.Take(&userModel, "user_name = ?", u.UserName).Error
	fmt.Println(dberr)
	if dberr != nil {
		// 没找到用户
		global.Log.Warn("用户名不存在")
		response.Fail(dberr, "用户名不存在!请重新输入。", c)
		return
	}

	isCheck := pwd.PassWordVerify(u.PassWord, userModel.Password)
	if !isCheck {
		global.Log.Warn("用户名密码错误")
		response.Fail(nil, "密码错误！请重新输入。", c)
		return
	}

	// 登陆成功， 生成token
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		NickName: userModel.NickName,
		Role:     int(userModel.Role),
		UserID:   userModel.ID,
		UserName: userModel.UserName,
	})
	if err != nil {
		global.Log.Error(err)
		response.Fail(nil, "token生成失败", c)
		return
	}
	response.Success(token, "success", c)
}
