package image

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zeropage/mukgoorm/setting"
)

func init() {
	dir := setting.GetDirectory()
	dir.Path = "../testdata"
}

func before() {
	MakeImageDir()
}

func after() {
	if f, _ := os.Stat(ImagePath()); f != nil {
		os.RemoveAll(ImagePath())
	}
}

func TestResize(t *testing.T) {
	before()
	defer after()

	fileName := "pic.jpg"
	Resize(path.Join(setting.GetDirectory().Path, fileName), 300)

	f, err := os.Stat(path.Join(ImagePath(), fileName))
	assert.NotNil(t, f)
	assert.Nil(t, err)
}
