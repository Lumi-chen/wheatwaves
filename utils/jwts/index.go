package jwts

import (
	"errors"
	"fmt"
	"time"
	"wheatwaves/global"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/google/uuid"
)

// jwt (json web token) 由三段落构成，每个段落用 . 连接，分别是 Header、Payload、Signature
// 内容会被base64加密，用base64解密就可以看到实际传输内容

// jwt里的payload数据
type JwtPayLoad struct {
	UserName string    `json:"user_name"` // 用户名
	NickName string    `json:"nick_name"` // 昵称
	Role     int       `json:"role"`      // 角色
	UserID   uuid.UUID `json:"user_id"`   // 用户id
}

type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims
}

var MySecret []byte // 密钥

// 创建Token
func GenToken(user JwtPayLoad) (string, error) {
	MySecret = []byte(global.Config.Jwy.Secret)
	claim := CustomClaims{
		user, // 用户信息
		jwt.StandardClaims{
			// 过期时间、签发人
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * time.Duration(global.Config.Jwy.Expires))), // 默认两小时过期
			Issuer:    global.Config.Jwy.Issuer,                                                     // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(MySecret)
}

func ParseToken(tokenStr string) (*CustomClaims, error) {
	MySecret = []byte(global.Config.Jwy.Secret)
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return MySecret, nil
		})
	if err != nil {
		global.Log.Error(fmt.Sprintf("token parse err: %s", err.Error()))
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// 使用前置，读取配置
// 读取配置文件
// core.InitConf()
// // 初始化日志
// global.Log = core.InitLogger()

// 生成token：
// token, err := jwts.GenToken(jwts.JwtPayLoad{
// 	UserID:   1,
// 	Role:     int(mtype.Admin),
// 	UserName: "fengfeng",
// 	NickName: "xxx",
// })
// fmt.Println(token, err)

// 解析
// clasime, clasimeErr := jwts.ParseToken(token)
// fmt.Println(clasime, clasimeErr)
