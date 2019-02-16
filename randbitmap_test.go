// Copyright 2019 Nikola Trubitsyn. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package randbitmap

import (
	"testing"
)

var fakeWriter FakeWriter

type FakeWriter struct {
}

func (w FakeWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}

func TestBlackImage(t *testing.T) {
	var height, width uint = 100, 100
	bitmap := NewBitmap(width, height, false)
	bitmap.SetGenerator(func(n int) int {
		return 0
	})
	if err := bitmap.Render(fakeWriter); err != nil {
		t.Fail()
	}

	for x := 0; x < int(width); x++ {
		for y := 0; y < int(height); y++ {
			r, g, b, _ := bitmap.image.At(x, y).RGBA()
			if !(r == 0 && g == 0 && b == 0) {
				t.FailNow()
			}
		}
	}
}

func TestWhiteImage(t *testing.T) {
	var height, width uint = 100, 100
	bitmap := NewBitmap(width, height, false)
	bitmap.SetGenerator(func(n int) int {
		return 1
	})
	if err := bitmap.Render(fakeWriter); err != nil {
		t.Fail()
	}

	for x := 0; x < int(width); x++ {
		for y := 0; y < int(height); y++ {
			r, g, b, _ := bitmap.image.At(x, y).RGBA()
			if !(r == 65535 && g == 65535 && b == 65535) {
				t.FailNow()
			}
		}
	}
}
