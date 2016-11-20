package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFileInfoAndPath(t *testing.T) {
	root := "../tmp/dat"
	result, err := getFileInfoAndPath(root)
	assert.Equal(t, err, nil)
	assert.NotZero(t, len(*result))
}

func TestGetFileInfoAndPathFail(t *testing.T) {
	root := "nodir"
	result, err := getFileInfoAndPath(root)
	assert.Error(t, err)
	assert.Zero(t, len(*result))
}
