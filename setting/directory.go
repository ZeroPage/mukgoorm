package setting

import (
	"path/filepath"
	"strings"
)

type directory struct {
	Path string
}

var dir *directory

func GetDirectory() *directory {
	return dir
}

func (d *directory) Valid(path string) bool {
	if strings.HasPrefix(filepath.ToSlash(path), filepath.ToSlash(d.Path)) {
		return true
	}
	return false
}

func init() {
	dir = &directory{}
}
