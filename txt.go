package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"strings"

	_ "golang.org/x/image/colornames"
	"golang.org/x/image/font"
	"golang.org/x/image/font/inconsolata"
	"golang.org/x/image/math/fixed"
)

const width, height, padding, lineHeight int = 212, 104, 20, 20

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	draw.Draw(img, img.Bounds(), &image.Uniform{color.Black}, image.ZP, draw.Src)
	drawText(img, padding, padding+8, "Hi this is a longer bit of text to test wrapping. Does it work Obi? Is this even longer?")

	if err := saveImage(img, "image.png"); err != nil {
		panic(err)
	}
}

func drawText(img draw.Image, x, y int, s string) {
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.White),
		Face: inconsolata.Bold8x16,
		Dot:  fixed.P(x, y),
	}

	lines := multiline(d, s)

	for li, line := range lines {
		d.Dot = fixed.P(x, y+(li*lineHeight))
		d.DrawString(line)
	}
}

func multiline(d *font.Drawer, s string) []string {
	words := strings.Fields(s)
	lineWidth := width - (padding * 2)
	lines := []string{""}
	li := 0

	for _, word := range words {
		if d.MeasureString(lines[li]+word) > fixed.I(lineWidth) {
			lines = append(lines, "")
			li++
		}
		lines[li] = lines[li] + word + " "
	}

	return lines
}

func saveImage(img draw.Image, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error saving image: %v", err)
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		return fmt.Errorf("error saving image: %v", err)
	}

	return nil
}
