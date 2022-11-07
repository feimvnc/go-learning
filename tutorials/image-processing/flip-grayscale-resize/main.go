// source https://go-recipes.dev/more-working-with-images-in-go-30b11ab2a9f0
// flip an image and save
// parrot.png - Photo by David Clode on Unsplash
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func main() {

	grid := load("parrot.png")

	gray := grayscale(grid)
	save("parrot_grayscale.png", gray)

	scale := resize(grid, 0.5)
	save("parrot_resize.png", scale)

	flip(grid)
	save("parrot_filpped.png", grid)
}

func flip(grid [][]color.Color) {
	for x := 0; x < len(grid); x++ {
		col := grid[x]
		for y := 0; y < len(col)/2; y++ { // only need top half pixels of y
			z := len(col) - y - 1           // get opposite position , skip middle(?)
			col[y], col[z] = col[z], col[y] // swap current and opposition position
		}
	}
}

func save(filePath string, grid [][]color.Color) {
	// create an image and set the pixels using the grid
	xlen, ylen := len(grid), len(grid[0]) // define width and height
	rect := image.Rect(0, 0, xlen, ylen)
	img := image.NewNRGBA(rect)
	for x := 0; x < xlen; x++ {
		for y := 0; y < ylen; y++ {
			img.Set(x, y, grid[x][y])
		}
	}
	// create a file and encode the image into it
	file, err := os.Create(filePath)
	if err != nil {
		log.Println("cannot create file:", err)
	}
	defer file.Close()
	png.Encode(file, img.SubImage(img.Rect))
}

func load(filePath string) (grid [][]color.Color) {
	//open the file and decode the contents into an image
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("cannot read file:", err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		log.Println("cannot decode file:", err)
	}
	// create and return a grid of pixels
	size := img.Bounds().Size()
	for i := 0; i < size.X; i++ {
		var y []color.Color
		for j := 0; j < size.Y; j++ {
			y = append(y, img.At(i, j)) // append to y
		}
		grid = append(grid, y) // append to x
	}
	return
}

func grayscale(grid [][]color.Color) (grayImg [][]color.Color) {
	xlen, ylen := len(grid), len(grid[0]) // get same dimensions as the original
	grayImg = make([][]color.Color, xlen)
	for i := 0; i < len(grayImg); i++ {
		grayImg[i] = make([]color.Color, ylen)
	}
	for x := 0; x < xlen; x++ {
		for y := 0; y < ylen; y++ {
			pix := grid[x][y].(color.NRGBA)
			// below is luminance of the pixel
			//gray := uint8(float64(pix.R)/3.0 + float64(pix.G)/3.0 + float64(pix.B)/3.0)
			// ITU-R BT.709 standard formula to convert color to grayscale
			// L = 0.2126 * R + 0.7152 * G + 0.0722 * B
			gray := uint8(0.2126*float64(pix.R) + 0.7152*float64(pix.G) + 0.0722*float64(pix.B))
			grayImg[x][y] = color.NRGBA{gray, gray, gray, pix.A}
		}
	}
	return
}

func resize(grid [][]color.Color, scale float64) (resized [][]color.Color) {
	// create new reized image
	xlen, ylen := int(float64(len(grid))*scale), int(float64(len(grid[0]))*scale)
	resized = make([][]color.Color, xlen)
	for i := 0; i < len(resized); i++ {
		resized[i] = make([]color.Color, ylen)

	}
	for x := 0; x < xlen; x++ {
		for y := 0; y < ylen; y++ {
			xp := int(math.Floor(float64(x) / scale))
			yp := int(math.Floor(float64(y) / scale))
			resized[x][y] = grid[xp][yp] //copy pixel from original to new image
		}
	}
	return
}
