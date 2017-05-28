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
	if val, ok := sessionVal.(int); !ok {
		return FAIL, ok
	} else {
		return Grant(val), ok
	}
}

func Name(g Grant) string {
	switch g {
	case READ_ONLY:
		return "Read only"
	case ADMIN:
		return "Admin"
	}
	return "Guest"
}
