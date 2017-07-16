package image

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJpg(t *testing.T) {
	res := signature("../tmp/dat/.images/2017-07-09194024_flow-dual-ring-1-jpeg.jpg")

	assert.Equal(t, res, "jpeg")
}

func TestPng(t *testing.T) {
	res := signature("../tmp/dat/2017-07-16172750_png_wolf_by_itsdura-d3cle9k.png")

	assert.Equal(t, res, "png")
}
