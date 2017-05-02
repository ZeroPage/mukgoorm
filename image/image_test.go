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
	imageDir := ImagePath()
	if f, _ := os.Stat(imageDir); f != nil {
		os.RemoveAll(imageDir)
	}
}

func TestResize(t *testing.T) {
	before()
	defer after()

	fileName := "pic.jpg"
	dir := setting.GetDirectory().Path
	println(fileName, dir)
	Resize(300, path.Join(dir, fileName))

	f, err := os.Stat(path.Join(ImagePath(), fileName))
	assert.NotNil(t, f)
	assert.Nil(t, err)
}
