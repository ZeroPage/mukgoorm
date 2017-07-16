package image

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zeropage/mukgoorm/setting"
)

func init() {
	dir := setting.GetDirectory()
	dir.Path = "../testdata"
}

func TestJpg(t *testing.T) {
	fileName := "pic.jpg"
	res := signature(path.Join(setting.GetDirectory().Path, fileName))

	assert.Equal(t, res, "jpeg")
}

func TestPng(t *testing.T) {
	fileName := "pic.png"
	res := signature(path.Join(setting.GetDirectory().Path, fileName))

	assert.Equal(t, res, "png")
}
