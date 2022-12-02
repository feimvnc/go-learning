package main

import (
	"fmt"
	"image/color"

	"github.com/skip2/go-qrcode"
)

func main() {
	err := qrcode.WriteFile("this is black-white", qrcode.Medium, 256, "bw-qr.png")
	if err != nil {
		fmt.Printf("could not create qrcode: %v", err)
	}

	// color.RGBA{255, 0, 0, 255}, r, g, b, a
	err = qrcode.WriteColorFile("this is colored", qrcode.Medium, 256, color.White, color.RGBA{255, 0, 0, 255}, "color-qr.png")
	if err != nil {
		fmt.Printf("could not create qrcode: %v", err)
	}

}
