package setting

import (
	"sync"
)

type directory struct {
	Path string
}

var dir *directory
var once sync.Once

func GetDirectory() *directory {
	once.Do(func() {
		dir = &directory{}
	})
	return dir
}
