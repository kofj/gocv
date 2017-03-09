package gocv

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/gif"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"os"

	"golang.org/x/image/bmp"
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/riff"
	"golang.org/x/image/tiff"
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

// Write image to graphics file
func Imwrite(img *image.Image, name, format string) (err error) {
	buf := new(bytes.Buffer)
	switch format {
	case "jpg", "jpeg":
		err = jpeg.Encode(buf, *img, nil)
	case "png":
		err = png.Encode(buf, *img)
	case "gif":
		err = gif.Encode(buf, *img, nil)
	case "bmp":
		err = bmp.Encode(buf, *img)
	case "tiff":
		err = tiff.Encode(buf, *img, nil)
	default:
		err = errors.New("not support")
	}
	if err != nil {
		return
	}

	w, _ := os.Create(fmt.Sprintf("%s.%s", name, format))
	fmt.Println(w.Write(buf.Bytes()))
	fmt.Println(len(buf.Bytes()))

	return
}
