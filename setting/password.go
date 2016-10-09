package setting

import (
	"sync"
)

type password struct {
	AdminPassword, ReadOnlyPassword string
}

var pw *password
var passwordLock sync.Once

func GetPassword() *password {
	passwordLock.Do(func() {
		pw = &password{}
	})
	return pw
}
