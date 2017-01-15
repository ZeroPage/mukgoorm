package handlers

import (
  "os"
	"testing"
  "archive/zip"
	"github.com/stretchr/testify/assert"
  "path/filepath"
  "strings"
)

func TestMakeZip(t *testing.T) {
  err := os.Mkdir("./test", 0777)
  file1, err := os.Create("test.txt")
  assert.Equal(t, err, nil)
  defer file1.Close()
  file1.WriteString("test")
  defer os.Remove("test.txt")

  filePath := "./test"
  fileName, err := makeZip(filePath)
  assert.Equal(t, err, nil)

  r, err := zip.OpenReader(fileName)
  assert.Equal(t, err, nil)

  num := 0
  filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error {
    _, name := filepath.Split(filePath)
    name += strings.TrimPrefix(path, filePath)
    assert.Equal(t, r.File[num].Name, name)
    file, err := os.Open(path)
    zfile, err := r.File[num].Open()

    state2 := archive/zip.FileInfo(r.File[num])
    state2, err := file.Stat()

    assert.Equal(t, state1.zfile.Name(), state2.IsDir())
    assert.Equal(t, zfile.Size(), state2.Size())
    num++
    return nil

  })
}
