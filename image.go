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

type Histogram [256]int

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
	w.Write(buf.Bytes())
	// fmt.Println(len(buf.Bytes()))

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
func Imhist(gray image.Gray) (hist *Histogram) {
	var (
		idx    int
		r      = gray.Bounds()
		width  = r.Dx()
		height = r.Dy()
	)
	hist = new(Histogram)
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			idx = i*height + j
			v := gray.Pix[idx]
			hist[v]++
		}
	}
	return
}

// Convert image to binary image, based on threshold
func Im2bw(img image.Gray, level uint8) (bw *image.Gray) {
	bw = image.NewGray(img.Bounds())
	if level == 0 || level == 255 {
		level = 255 / 2
	}
	for i := 0; i < len(img.Pix); i++ {
		if img.Pix[i] > level {
			bw.Pix[i] = 255
		} else {
			bw.Pix[i] = 0
		}
	}
	return
}

// Global histogram threshold using Otsu's method
func Otsuthresh(hist *Histogram, size int) (threshold uint8) {
	// normalize histogram
	// 直方图归一化
	var nHist [256]float64
	for i := 0; i < 256; i++ {
		nHist[i] = float64(hist[i]) / float64(size)
	}

	// average pixel value
	// 整幅图像的平均灰度
	var avgVal float64
	for i := 0; i < 256; i++ {
		avgVal += float64(i) * nHist[i]
	}

	var maxVariance float64
	var w, u float64
	for i := 0; i < 256; i++ {
		// 假设当前灰度i为阈值, 0~i 灰度的像素(假设像素值在此范围的像素叫做前景像素) 所占整幅图像的比例
		w += nHist[i]
		// 灰度i 之前的像素(0~i)的平均灰度值： 前景像素的平均灰度值
		u += float64(i) * nHist[i]

		t := avgVal*w - u
		variance := t * t / (w * (1 - w))
		if variance > maxVariance {
			maxVariance = variance
			threshold = uint8(i)
		}
	}
	return
}
