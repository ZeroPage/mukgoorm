package image

import (
	"bufio"
	"bytes"
	"os"
)

func signature(path string) string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	data := scanner.Bytes()
	if bytes.HasPrefix(data, []byte("\xFF\xD8\xFF")) {
		return "jpeg"
	} else if bytes.HasPrefix(data, []byte("\x89\x50\x4E\x47\x0D\x0A\x1A\x0A")) {
		return "png"
	} else if bytes.HasPrefix(data, []byte("\x89PNG")) {
		return "png"
	}
	return ""
}
