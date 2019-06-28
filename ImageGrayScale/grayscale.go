/*
This file converts the image to gray scale image

Source: https://riptutorial.com/go/example/31693/convert-color-image-to-grayscale
*/

package main

import (
	"image"
	"log"
	"net/http"
	"os"

	"image/jpeg"
	"image/png"
)

func main() {
	// Load image from remote through http
	// The Go gopher was designed by Renee French. (http://reneefrench.blogspot.com/)
	// Images are available under the Creative Commons 3.0 Attributions license.
	resp, err := http.Get("http://golang.org/doc/gopher/fiveyears.jpg")
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Decode image to JPEG
	img, _, err := image.Decode(resp.Body)
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	log.Printf("Image type: %T", img)

	file, err := os.Create("ImageGrayScale/fiveyears.jpg")
	defer file.Close()
	if err != nil {
		log.Fatal("Failed to create file, Error: ", err.Error())
	}
	defer file.Close()

	if err := jpeg.Encode(file, img, &jpeg.Options{Quality: 100}); err != nil {
		log.Fatal(err)
	}

	// Converting image to grayscale
	grayImg := image.NewGray(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			grayImg.Set(x, y, img.At(x, y))
		}
	}

	// Working with grayscale image, e.g. convert to png
	f, err := os.Create("ImageGrayScale/fiveyears_gray.png")
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	defer f.Close()

	if err := png.Encode(f, grayImg); err != nil {
		log.Fatal(err)
	}
}
