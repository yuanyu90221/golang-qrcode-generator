package main

import (
	"fmt"
	"image/color"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	// generate colorful qrcode
	err := qrcode.WriteColorFile("hello this is qrcode in golang", qrcode.Medium, 256, color.RGBA{0, 0, 255, 255}, color.White, "second_file.png")
	if err != nil {
		fmt.Printf("Error for create qrcode: %v", err)
	}
}
