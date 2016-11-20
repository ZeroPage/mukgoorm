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

func FromPassword(password string) Grant {
	sharePassword := setting.GetPassword()

	switch password {
	case sharePassword.AdminPassword:
		return ADMIN
	case sharePassword.ReadOnlyPassword:
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
