package path

import (
	"os"
	"path/filepath"
)

type OPTION int

const (
	NONE     OPTION = 0
	SKIP_DIR        = 1
)

func PathInfoWithDirFrom(root string) (*[]FilePathInfo, error) {
	return walk(root, NONE)
}

func PathInfoFrom(root string) (*[]FilePathInfo, error) {
	return walk(root, SKIP_DIR)
}

func walk(root string, options OPTION) (*[]FilePathInfo, error) {
	files := []FilePathInfo{}
	err := filepath.Walk(root, filepath.WalkFunc(func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path == root {
			return nil
		}

		files = append(files, FilePathInfo{info, path})
		if info.IsDir() && (options&SKIP_DIR != 0) {
			return filepath.SkipDir
		}
		return err
	}))
	return &files, err
}

type FilePathInfo struct {
	File os.FileInfo
	Path string
}
