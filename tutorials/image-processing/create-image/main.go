// source: https://go-recipes.dev/manipulating-2d-images-in-go-96106dd1b93f
// create an new random image

package main

import (
	"crypto/rand"
	"image"
	"image/png"
	"log"
	"os"
)

func main() {
	rect := image.Rect(0, 0, 100, 100)
	img := createRandomImage(rect)
	save("random.png", img)
}

func createRandomImage(rect image.Rectangle) (created *image.NRGBA) {
	pix := make([]uint8, rect.Dx()*rect.Dy()*4)
	rand.Read(pix)
	created = &image.NRGBA{
		Pix: pix,
		// Stride: rect.Dx() * 4,
		Stride: rect.Dx() * 4,
		Rect:   rect,
	}
	return
}

func save(filePath string, img *image.NRGBA) {
	imgFile, err := os.Create(filePath)
	defer imgFile.Close()
	if err != nil {
		log.Panicln("cannot create file:", err)
	}

	png.Encode(imgFile, img.SubImage(img.Rect))
}
