package image

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/nfnt/resize"
	"github.com/zeropage/mukgoorm/setting"
)

const (
	JPG_EXTEND  = "jpg"
	JPEG_EXTEND = "jpeg"
	PNG_EXTEND  = "png"
)

func FileExtend(filename string) string {
	s := strings.Split(filename, ".")
	return s[len(s)-1]
}

func IsImage(filename string) bool {
	extend := FileExtend(filename)
	if extend == JPG_EXTEND || extend == PNG_EXTEND {
		return true
	}
	return false
}

func Resize(imagePath string, size uint) {
	t := signature(imagePath)
	if t != JPEG_EXTEND && t != PNG_EXTEND {
		return
	}

	d, err := ioutil.ReadFile(imagePath)
	if err != nil {
		panic(err)
	}
	Compress(size, imagePath, d)
}

func Compress(size uint, imagePath string, data []byte) {
	var img image.Image
	var err error
	if extend := FileExtend(imagePath); extend == PNG_EXTEND {
		img, err = png.Decode(bytes.NewReader(data))
	} else {
		img, _, err = image.Decode(bytes.NewReader(data))
	}
	if err != nil {
		panic(err)
	}
	newImg := resize.Resize(size, size, img, resize.Lanczos3)

	s := strings.Split(imagePath, "/")
	s = strings.Split(s[len(s)-1], ".")
	name := s[0] + "." + JPG_EXTEND
	dir := path.Join(ImagePath(), name)
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
