package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"net/http"
	"time"
)

// Измените программу lissajous так, чтобы она генерировала
// изображения разных цветов, добавляя в палитру palette больше значений, а затем
// выводя их путем изменения третьего аргумента функции SetСolorlndex некоторым
// нетривиальным способом.

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		lissajous(w)
	})

	http.ListenAndServe("localhost:9999", nil)
}

var palette = []color.Color{
	color.Black,
	color.RGBA{255, 0, 0, 255},   // красный
	color.RGBA{0, 255, 0, 255},   // зеленый
	color.RGBA{0, 0, 255, 255},   // синий
	color.RGBA{255, 255, 0, 255}, // желтый
	color.RGBA{255, 0, 255, 255}, // фиолетовый
	color.RGBA{0, 255, 255, 255}, // голубой
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // Количество полных колебаний x
		res     = 0.001 // Угловое разрешение
		size    = 100   // Канва изображения охватывает [size..+size]
		nframes = 64    // Количество кадров анимации
		delay   = 8     // Задержка между кадрами (единица - 10мс)
	)

	rand.New(rand.NewSource(time.Now().UnixNano()))

	freq := rand.Float64() * 3.0        // частота обновления
	anim := gif.GIF{LoopCount: nframes} // создание анимации
	phase := 0.0                        // фаза

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			// Выбор цвета из палитры по индексу, зависящему от t и i
			colorIndex := uint8(1 + int((t*10)+float64(i))%(len(palette)-1))

			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}
