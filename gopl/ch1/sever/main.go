package main

import (
	"fmt"
	"log"
	"net/http"
    "image"
    "image/color"
    "image/gif"
    "io"
    "math"
    "math/rand"
)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/lissajous", func(w http.ResponseWriter, r *http.Request) {
            fmt.Println("200 OK - ", r.URL.Path)
            lissajous(w)
    })

    fmt.Println("Running server at localhot:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("200 OK - %s - ", r.URL.Path)
	fmt.Println("Hello World!")
}

var palette = []color.Color{
        color.Black,
        color.RGBA{0x00, 0xFF, 0x00, 0xFF},
        color.RGBA{0x00, 0x00, 0xFF, 0xFF},
        color.RGBA{0xFF, 0xCC, 0xFF, 0xFF},
        color.RGBA{0x0F, 0xCC, 0x0F, 0xFF},
        color.RGBA{0x00, 0x00, 0xFF, 0xFF},
}

func lissajous(out io.Writer) {
    const (
        cycles  = 5     // número de revoluções completas do oscilador x
        res     = 0.001 // resolução angular
        size    = 100   // canvas da imagem cobre de [-size..+size]
        nframes = 64    // número de quadros da animação
        delay   = 8     // tempo entre quadros em unidades de 10ms
    )
    freq := rand.Float64() * 3.0 // frequência relativa do oscilador y
    anim := gif.GIF{LoopCount: nframes}
    phase := 0.0 // diferença de fase

    for i := 0; i < nframes; i++ {
        rect := image.Rect(0, 0, 2*size+1, 2*size+1)
        img := image.NewPaletted(rect, palette)
        color := 0
        for t := 0.0; t < cycles*2*math.Pi; t += res {
            x := math.Sin(t)
            y := math.Sin(t*freq + phase)
            img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(color))
            color++
            if color >= len(palette) {
                color = 0
            }
        }
        phase += 0.1
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }
    gif.EncodeAll(out, &anim)
}
