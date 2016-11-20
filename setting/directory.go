package setting

import (
	"path/filepath"
	"strings"
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

func (d *directory) ValidDir(path string) bool {
	if strings.HasPrefix(filepath.ToSlash(path), filepath.ToSlash(d.Path)) {
		return true
	}
	return false
}
