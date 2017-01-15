package main

import (
  "archive/zip"
  "os"
	"path/filepath"
  "io"
  "strings"
)

func main(){
  os.Mkdir("./test", 0777)
}

func makeZip(foldername string) (string, error) {
	newfile, err := os.Create(foldername + ".zip")
	if err != nil {
		return "", err
	}
	defer newfile.Close()

	zipit := zip.NewWriter(newfile)
	defer zipit.Close()

	filepath.Walk(foldername, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
    _, name := filepath.Split(foldername)
    header.Name = name
    header.Name += strings.TrimPrefix(path, foldername)


		if info.IsDir() {

		} else {
			header.Method = zip.Deflate
		}
		writer, err := zipit.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		zipfile, err := os.Open(path)
		defer zipfile.Close()
		if err != nil {
			return err
		}

		_, err = io.Copy(writer, zipfile)
		return err

	})
	return foldername + ".zip", err
}
