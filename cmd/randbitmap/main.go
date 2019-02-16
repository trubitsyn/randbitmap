// Copyright 2019 Nikola Trubitsyn. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/trubitsyn/randbitmap"
	"math/rand"
	"os"
)

var (
	size      = flag.Uint("size", 100, "Bitmap size")
	useColor  = flag.Bool("color", false, "Color bitmap")
	filename  = flag.String("file", "random.png", "Output file")
	overwrite = flag.Bool("overwrite", false, "Overwrite existing file")
)

func main() {
	flag.Parse()
	if *size == 0 {
		die(errors.New("image size must be greater than zero"))
	}
	bitmap := randbitmap.NewBitmap(*size, *size, *useColor)
	bitmap.SetGenerator(func(n int) int {
		return rand.Intn(n)
	})
	if err := renderBitmap(bitmap, *filename); err != nil {
		die(err)
	}
}

func die(err error) {
	fmt.Println("Error:", err.Error())
	os.Exit(1)
}

func renderBitmap(bitmap *randbitmap.Bitmap, filename string) error {
	f, err := openFile(filename)
	if err != nil {
		return err
	}
	return bitmap.Render(f)
}

func openFile(filename string) (*os.File, error) {
	flags := os.O_RDWR | os.O_CREATE | os.O_EXCL
	if *overwrite {
		flags &^= os.O_EXCL
	}
	f, err := os.OpenFile(filename, flags, 0666)
	if err != nil {
		return nil, err
	}
	return f, nil
}
