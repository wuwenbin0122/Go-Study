// lissajous产生随机利萨如图形
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var palette = []color.Color{color.White, color.Black,
	color.RGBA{255, 0, 0, 0xff}, color.RGBA{0, 255, 0, 0xff},
	color.RGBA{0, 0, 255, 0xff}, color.RGBA{192, 192, 192, 0xff},
	color.RGBA{255, 255, 0, 0xff}, color.RGBA{255, 0, 255, 0xff}}

const (
	whiteIndex  = 0 // first color in palette
	blackIndex  = 1 // next color in palette
	redIndex    = 2
	greenIndex  = 3
	blueIndex   = 4
	greyIndex   = 5
	yellowIndex = 6
	ceriseIndex = 7
)

func main() {
	rand.NewSource(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     //完整的x振荡器变化的个数
		res     = 0.001 //角度分辨率
		size    = 100   //画布包含[-size..+size]
		nframes = 64    //动画中的帧数
		delay   = 8     //以10ms为单位的帧间延迟
	)
	freq := rand.Float64() * 3.0 //y振荡器的相对频率
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 //phase difference
	c := 0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(c))
		}
		if c >= len(palette) {
			c = 0
		}
		c++
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) //忽略编码错误
}
