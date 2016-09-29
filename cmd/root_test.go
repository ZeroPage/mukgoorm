package cmd

import (
	"strings"
	"testing"

	"../setting"

	"github.com/stretchr/testify/assert"
)

func TestLongFlagSetDirectory(t *testing.T) {
	input := "--dir tmp/dat"
	RootCmd.SetArgs(strings.Split(input, " "))

	err := RootCmd.Execute()
	assert.Equal(t, err, nil)

	setting := setting.GetDirectory()
	assert.Equal(t, setting.Path, "tmp/dat")
}

func TestShortFlagSetDirectory(t *testing.T) {
	input := "-D tmp/dat"
	RootCmd.SetArgs(strings.Split(input, " "))

	err := RootCmd.Execute()
	assert.Equal(t, err, nil)

	setting := setting.GetDirectory()
	assert.Equal(t, setting.Path, "tmp/dat")
}

func TestInvalidPath(t *testing.T) {
	input := "--dir nodir"
	RootCmd.SetArgs(strings.Split(input, " "))

	err := RootCmd.Execute()
	assert.Equal(t, err, nil)

	// TODO(rabierre): should validate?
	setting := setting.GetDirectory()
	assert.Equal(t, setting.Path, "nodir")
}

func TestInvalidFlag(t *testing.T) {
	input := "--no-flag tmp/dat"
	RootCmd.SetArgs(strings.Split(input, " "))

	err := RootCmd.Execute()
	assert.NotEqual(t, err, nil)
}
