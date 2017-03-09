package gocv

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/draw"
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
func Imread(name string) (img image.Image, format string, err error) {
	var f *os.File

	f, err = os.Open(name)
	if err != nil {
		return
	}
	defer f.Close()

	img, format, err = image.Decode(f)

	return
}

// Write image to graphics file
func Imwrite(img image.Image, name, format string) (err error) {
	buf := new(bytes.Buffer)
	switch format {
	case "jpg", "jpeg":
		err = jpeg.Encode(buf, img, nil)
	case "png":
		err = png.Encode(buf, img)
	case "gif":
		err = gif.Encode(buf, img, nil)
	case "bmp":
		err = bmp.Encode(buf, img)
	case "tiff":
		err = tiff.Encode(buf, img, nil)
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

// Convert to gray image
func Im2gray(img image.Image) (gray *image.Gray) {
	r := img.Bounds()
	gray = image.NewGray(r)
	draw.Draw(gray, r, img, img.Bounds().Min, draw.Src)
	return
}

// Histogram of image data
func Imhist(gray image.Gray) (hist []int) {
	var (
		idx    int
		r      = gray.Bounds()
		width  = r.Dx()
		height = r.Dy()
	)
	hist = make([]int, 256)
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			idx = i*height + j
			v := gray.Pix[idx]
			hist[v]++
		}
	}
	return
}
