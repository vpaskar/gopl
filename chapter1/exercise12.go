// starting form 1.4
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

var palette = []color.Color{color.White, color.RGBA{0, 0xff, 0, 0xff}, color.RGBA{0xff, 0, 0, 0xff}, color.RGBA{0, 0, 0xff, 0xff}, color.RGBA{0xff, 0xff, 0xff, 0xff}}

const (
	whiteIndex = 0
	blackindex = 3
)

func exercise15(out io.Writer, cycles float64) {
	const (
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq, anim, phase := rand.Float64(), gif.GIF{LoopCount: nframes}, 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x, y := math.Sin(t), math.Sin(t*freq+phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackindex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func exercise16(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq, anim, phase := rand.Float64(), gif.GIF{LoopCount: nframes}, 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x, y := math.Sin(t), math.Sin(t*freq+phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackindex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
