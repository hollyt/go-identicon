package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

// Set up command line flag variables
var hashString string

func init() {
	flag.StringVar(&hashString, "hs", "Hello", "String to hash")
}

// Create the identicon image
type identicon struct {
	size        int
	rows        int
	cols        int
	hashedBytes []byte
}

func New(hashedBytes []byte) *identicon {
	return &identicon{
		size:        70,
		rows:        5,
		cols:        5,
		hashedBytes: hashedBytes,
	}
}

func (id *identicon) createIdenticon() {
	background := color.RGBA{255, 255, 255, 255}
	colors := id.hashedBytes[0:3]
	r, g, b := colors[0], colors[1], colors[2]
	pixelColor := color.RGBA{r, g, b, 1}

	img := image.NewPaletted(image.Rect(0, 0, 1000, 1000), color.Palette{background, pixelColor})
	for i := 100; i < 1000; i += 1 {
		img.Pix[i] = r
		img.Pix[i+1] = g
		img.Pix[i+2] = b
		img.Pix[i+3] = 1
	}
	outputFile, err := os.Create("test.png")
	if err != nil {
		fmt.Println("Error creating .png")
		return
	}
	png.Encode(outputFile, img)
}

func main() {

	// Get string to hash
	flag.Parse()

	hashed := md5.New()
	hashed.Write([]byte(hashString))
	hashedBytes := hashed.Sum(nil)

	id := New(hashedBytes)
	id.createIdenticon()

	// Testing
	fmt.Println("id: ", id)
	fmt.Println("The string to hash: ", hashString)
	fmt.Printf("Hash in hex: %x\n", hashedBytes)
	fmt.Println("Hash in bytes: ", hashedBytes)
}
