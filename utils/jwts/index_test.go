package jwts

import (
	"testing"
	"wheatwaves/core"
	"wheatwaves/global"
)

func TestJwt(t *testing.T) {
	core.InitConf()
	global.Log = core.InitLogger()

}
