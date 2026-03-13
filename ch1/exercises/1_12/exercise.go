package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

// This a server lissajous implementation
// Functions are
//	1. Create a web server minimal echo and counter
// 	2. Read query params to dynamically configure lissajous on web browser

var counter int
var mu sync.Mutex

var pallete = []color.Color{color.White, color.Black, color.RGBA{0x00, 0xFF, 0x00, 0xFF}, color.RGBA{0xFF, 0x00, 0x00, 0xFF}, color.RGBA{0xFF, 0xFF, 0x00, 0xFF}}

const (
	whiteIndex = iota
	blackIndex
	greenIndex
	redIndex
	brownIndex
)

func main() {
	// lissajous(w)
	http.HandleFunc("/", handler)

	counterHandler := func(w http.ResponseWriter, request *http.Request) {
		mu.Lock()
		fmt.Fprintf(w, "\n%d\n", counter)
		mu.Unlock()

	}
	http.HandleFunc("/count", counterHandler)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func handler(w http.ResponseWriter, request *http.Request) {
	mu.Lock()
	counter++
	mu.Unlock()

	defaultCycleCount := 5
	cyclesParam, err := strconv.ParseFloat(request.URL.Query().Get("cycles"), 64)
	if err != nil || cyclesParam <= 0 {
		lissajous(w, float64(defaultCycleCount))
	} else {
		lissajous(w, cyclesParam)
	}

}

func lissajous(out io.Writer, cycles float64) {
	const (
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
			//Function [i%5] makes the gif have all the values as per frame
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(i%5))

		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}
