// Lissajous generates GIF animation of randam Lissajous figures
package main

// Change the Lissajous Program's color palette to green on blakc for added authenticty

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var pallete = []color.Color{color.White, color.Black, color.RGBA{0x00, 0xFF, 0x00, 0xFF}}

const (
	whiteIndex = iota
	blackIndex
	greenIndex
)

func main() {
	lissajous(os.Stdout)

}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes,
		BackgroundIndex: blackIndex}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, pallete)
		//Color the rectangle
		for x := img.Rect.Min.X; x < img.Rect.Max.X; x++ {
			for y := img.Rect.Min.Y; y < img.Rect.Max.Y; y++ {
				img.SetColorIndex(x, y, blackIndex)
			}
		}

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)

		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}
