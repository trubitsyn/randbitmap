// Copyright 2019 Nikola Trubitsyn. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package randbitmap

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
)

type generator func(n int) int

type Bitmap struct {
	width     uint
	height    uint
	hasColor  bool
	generator generator
	image     *image.NRGBA
}

func NewBitmap(width uint, height uint, color bool) *Bitmap {
	return &Bitmap{
		width:    width,
		height:   height,
		hasColor: color,
	}
}

func (bitmap *Bitmap) SetGenerator(generator generator) {
	bitmap.generator = generator
}

func (bitmap *Bitmap) Render(w io.Writer) error {
	bitmap.image = image.NewNRGBA(image.Rect(0, 0, int(bitmap.width), int(bitmap.height)))
	draw.Draw(bitmap.image, bitmap.image.Bounds(), &image.Uniform{C: color.White}, image.ZP, draw.Src)
	colorFunc := getColorFunc(bitmap.hasColor, bitmap.generator)
	drawPixels(bitmap.image, colorFunc)
	if err := png.Encode(w, bitmap.image); err != nil {
		return err
	}
	return nil
}

func drawPixels(img *image.NRGBA, colorFunc func() (uint8, uint8, uint8)) {
	size := img.Rect.Size()
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			r, g, b := colorFunc()
			img.Set(x, y, color.NRGBA{R: r, G: g, B: b, A: 255})
		}
	}
}

func getColorFunc(color bool, generator generator) func() (uint8, uint8, uint8) {
	if color {
		return func() (uint8, uint8, uint8) {
			return colorPixel(generator)
		}
	} else {
		return func() (uint8, uint8, uint8) {
			return bwPixel(generator)
		}
	}
}

func bwPixel(generator generator) (uint8, uint8, uint8) {
	i := generator(2)
	if i == 0 {
		return 0, 0, 0
	}
	return 255, 255, 255
}

func colorPixel(generator generator) (uint8, uint8, uint8) {
	r := generator(255)
	g := generator(255)
	b := generator(255)
	return uint8(r), uint8(g), uint8(b)
}
