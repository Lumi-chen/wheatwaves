package pwd

import "golang.org/x/crypto/bcrypt"

// 密码加密
func PasswordHash(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

// 密码验证
func PassWordVerify(pwd, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}

// 检查密码是否一致
func CheckPwd(pwd, checkPwd string) bool {
	pwdCost, _ := bcrypt.Cost([]byte(pwd))
	checkCost, _ := bcrypt.Cost([]byte(checkPwd))
	return pwdCost == checkCost
	// err := bcrypt.CompareHashAndPassword([]byte(pwd), []byte(checkPwd))
	// return err == nil
}
