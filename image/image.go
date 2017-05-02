package image

import (
	"bytes"
	"image"
	"image/jpeg"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/nfnt/resize"
	"github.com/zeropage/mukgoorm/setting"
)

func IsImage(filename string) bool {
	s := strings.Split(filename, ".")
	extend := s[len(s)-1]
	if extend == "jpg" || extend == "png" {
		return true
	}
	return false
}

func Resize(size uint, imagePath string) {
	d, _ := ioutil.ReadFile(imagePath)
	Compress(size, imagePath, d)
}

func Compress(size uint, imagePath string, data []byte) {
	img, _, err := image.Decode(bytes.NewReader(data))
	newImg := resize.Resize(size, size, img, resize.Lanczos3)

	dir := ImagePath()
	s := strings.Split(imagePath, "/")
	name := s[len(s)-1]
	dir = path.Join(dir, name)
	out, err := os.Create(dir)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	jpeg.Encode(out, newImg, nil)
}

func ImagePath() string {
	return path.Join(setting.GetDirectory().Path, ".images")
}

func MakeImageDir() {
	imageDir := ImagePath()
	if f, err := os.Stat(imageDir); f == nil {
		err = os.Mkdir(imageDir, 0770)
		if err != nil {
			panic(err)
		}
	}
}
