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

func FromSession(sessionVal interface{}) Grant {
	val, ok := sessionVal.(int)
	if !ok {
		return FAIL
	}
	return Grant(val)
}

func AuthorityExist(grant Grant) (bool, error) {
	switch grant {
	case ADMIN, READ_ONLY:
		return true, nil
	default:
		return false, nil
	}
}
