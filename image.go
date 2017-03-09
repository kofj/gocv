package gocv

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/riff"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/vp8"
	_ "golang.org/x/image/vp8l"
	_ "golang.org/x/image/webp"
)

// Read image from graphics file
func Imread(name string) (img *image.Image, format string, err error) {
	var f *os.File
	img = new(image.Image)

	f, err = os.Open(name)
	if err != nil {
		return
	}
	defer f.Close()

	*img, format, err = image.Decode(f)

	return
}
