package day_14

import (
	"aoc24/bing"
	"aoc24/color"
	image2 "aoc24/image"
	bytes2 "bytes"
	"image"
	color2 "image/color"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
)

type Pattern = [][]float64

func rgbAt(img image.Image, x, y int) (uint8, uint8, uint8) {
	r, g, b, _ := img.At(x, y).RGBA()
	return uint8(r >> 8), uint8(g >> 8), uint8(b >> 8)
}

func createPattern(img image.Image, width, height, _ int) Pattern {
	resized := image2.Resize(img, width, height)
	rows := make([][]float64, height)
	for y := range height {
		row := make([]float64, width)
		for x := range width {
			r, g, b := rgbAt(resized, x, y)
			_, _, l := color.ToHSL(r, g, b)
			if l > 0.5 {
				row[x] = 1.0
			} else {
				row[x] = 0.0
			}
		}
		rows[y] = row
	}
	return rows
}

func averagePattern(patterns []Pattern) Pattern {
	height := len(patterns[0])
	width := len(patterns[0][0])

	rows := make([][]float64, height)
	for y := range height {
		row := make([]float64, width)
		for x := range width {
			value := 0.0
			for _, pattern := range patterns {
				value += pattern[y][x]
			}
			row[x] = value / float64(len(patterns))
		}
		rows[y] = row
	}
	return rows
}

func renderPattern(pattern Pattern) image.Image {
	height := len(pattern)
	width := len(pattern[0])
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := range pattern {
		for x := range pattern[y] {
			value := pattern[y][x] * 255
			img.Set(x, y, color2.RGBA{
				R: uint8(value),
				G: uint8(value),
				B: uint8(value),
				A: uint8(255),
			})
			if pattern[y][x] > 0.75 {
			}
		}
	}
	return img
}

func loadPattern(filename string) Pattern {
	file, _ := os.Open(filename)
	img, _ := png.Decode(file)
	width, height := img.Bounds().Dx(), img.Bounds().Dy()

	rows := make([][]float64, height)
	for y := range height {
		row := make([]float64, width)
		for x := range width {
			r, _, _ := rgbAt(img, x, y)
			row[x] = float64(r) / 255
		}
		rows[y] = row
	}
	return rows
}

func loadKey() string {
	file, _ := os.Open("./key.txt")
	bytes, _ := io.ReadAll(file)
	return string(bytes)
}

func trainPattern(output string) {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("jpeg", "jpg", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)

	key := loadKey()
	searchClient := bing.CreateImageSearchClient(key)
	searchResults := searchClient.Search("Christmas Tree")

	var patterns = make([]Pattern, 0)
	for _, result := range searchResults {
		resp, _ := http.Get(result.Thumbnail.Url)
		bytes, _ := io.ReadAll(resp.Body)
		buffer := bytes2.NewBuffer(bytes)
		decoded, _, _ := image.Decode(buffer)
		patterns = append(patterns, createPattern(decoded, 101, 103, 120))
	}
	average := averagePattern(patterns)
	rendered := renderPattern(average)
	file, err := os.OpenFile(output, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	err = png.Encode(file, rendered)
	if err != nil {
		panic(err)
	}
	_ = file.Close()
}
