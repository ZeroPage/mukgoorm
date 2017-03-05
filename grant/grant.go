package grant

import (
	"github.com/zeropage/mukgoorm/setting"
)

type Grant int

const (
	FAIL Grant = iota
	ADMIN
	READ_ONLY
)

func FromPassword(pwd string) Grant {
	shared := setting.GetPassword()

	switch pwd {
	case shared.AdminPwd:
		return ADMIN
	case shared.ROnlyPwd:
		return READ_ONLY
	default:
		return FAIL
	}
}

func FromSession(sessionVal interface{}) (Grant, bool) {
	val, ok := sessionVal.(int)
	if !ok {
		return FAIL, false
	}
	return Grant(val), true
}
