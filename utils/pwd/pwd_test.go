package pwd

import (
	"testing"
)

func TestA(t *testing.T) {
	t.Log("testA")
}

func TestPwdHash(t *testing.T) {
	pwd := "admin1234"
	hash, _ := PasswordHash(pwd)
	t.Log("密码：", pwd)
	t.Log("生成Hash : ", hash)

	match := PassWordVerify(pwd, hash)
	t.Log("验证结果", match)
}

func TestPwdVerify(t *testing.T) {
	pwd := "admin1234"
	hash := "$2a$10$Eqfb7IGJewkFJadmHMp8G.vmUt2zRZ03mw321fhLKSVvvqq8fsNP6"
	t.Log("密码：", pwd)
	t.Log("生成Hash : ", hash)

	match := PassWordVerify(pwd, hash)
	t.Log("验证结果", match)

}

func TestCheck(t *testing.T) {
	pwd := "$2a$10$xM1HLFi5YZ2sSXGMJ2yfQOkLep0PqCYc2aTAd2AKsF51r1ZTEfgeG"
	check := "$2a$10$Eqfb7IGJewkFJadmHMp8G.vmUt2zRZ03mw321fhLKSVvvqq8fsNP6"
	t.Log("pwd", pwd)
	t.Log("check : ", check)
	match := CheckPwd(pwd, check)
	t.Log("res:", match)
}
