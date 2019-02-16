# randbitmap [![Build Status](https://travis-ci.com/trubitsyn/randbitmap.svg?branch=master)](https://travis-ci.com/trubitsyn/randbitmap) [![GoDoc](https://godoc.org/github.com/trubitsyn/randbitmap?status.svg)](https://godoc.org/github.com/trubitsyn/randbitmap) [![Go Report Card](https://goreportcard.com/badge/github.com/trubitsyn/randbitmap)](https://goreportcard.com/report/github.com/trubitsyn/randbitmap)
Generate random bitmaps.

![Random black & white bitmap](bitmaps/bw.png)
![Random color bitmap](bitmaps/color.png)

## Installation
`go get -u github.com/trubitsyn/randbitmap/cmd/randbitmap`

## Usage
Execute `randbitmap` to save black & white 100x100 image to `random.png`.

```
Usage of randbitmap:
  -color
        Color bitmap
  -file string
        Output file (default "random.png")
  -overwrite
        Overwrite existing file
  -size uint
        Bitmap size (default 100)
```

### Defaults
`randbitmap -size=100 -color=false -file=random.png -overwrite=false`

## Library
### Installation
`go get -u github.com/trubitsyn/randbitmap`

### Usage
<pre>
package main

import (
	"flag"
	"fmt"
	"github.com/trubitsyn/randbitmap"
	"math/rand"
	"os"
)

var (
	width    = flag.Uint("width", 100, "Bitmap size")
	height   = flag.Uint("height", 100, "Bitmap size")
	color    = flag.Bool("color", false, "Color bitmap")
	filename = flag.String("file", "random.png", "Output file")
)

func main() {
	flag.Parse()
	bitmap := randbitmap.NewBitmap(*width, *height, *color)
	bitmap.SetGenerator(func(n int) int {
		return rand.Intn(n)
	})
	f, err := os.Create(*filename)
	if err != nil {
		fmt.Println(err)
	}
	if err := bitmap.Render(f); err != nil {
		fmt.Println(err)
	}
}
</pre>

## Testing
```
go get -t github.com/trubitsyn/randbitmap
go test github.com/trubitsyn/randbitmap
```

## LICENSE
```
Copyright 2019 Nikola Trubitsyn

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```