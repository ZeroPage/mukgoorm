package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zeropage/mukgoorm/setting"
)

func TestSearch(t *testing.T) {
	setting.GetDirectory().Path = "../testdata"

	res := search("")
	assert.Equal(t, len(res), 0)

	res = search("A")
	assert.Equal(t, len(res), 1)
}
