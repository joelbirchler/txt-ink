package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	_ "golang.org/x/image/colornames"
	"golang.org/x/image/font"
	"golang.org/x/image/font/inconsolata"
	"golang.org/x/image/math/fixed"
)

func main() {
	w, h := 212, 104
	img := image.NewRGBA(image.Rect(0, 0, w, h))

	draw.Draw(img, img.Bounds(), &image.Uniform{color.Black}, image.ZP, draw.Src)
	addText(img, 40, 20, "Hi")

	if err := saveImage(img, "image.png"); err != nil {
		panic(err)
	}
}

func addText(img draw.Image, x, y int, s string) {
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.White),
		Face: inconsolata.Bold8x16,
		Dot:  fixed.P(x, y),
	}

	d.DrawString(s)
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
