package path

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathInfoFrom(t *testing.T) {
	root := "../tmp/dat"
	result, err := PathInfoFrom(root)
	assert.Equal(t, err, nil)
	assert.NotZero(t, len(*result))
}

func TestPathInfoFromFail(t *testing.T) {
	root := "nodir"
	result, err := PathInfoFrom(root)
	assert.Error(t, err)
	assert.Zero(t, len(*result))
}
