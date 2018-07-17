// From: https://stackoverflow.com/questions/25271654/how-to-get-image-resolution-from-url-in-golang
package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://i.imgur.com/Peq1U1u.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	m, f, err := image.Decode(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("format:", f)
	g := m.Bounds()

	// Get height and width
	height := g.Dy()
	width := g.Dx()

	// The resolution is height x width
	resolution := height * width

	// Print results
	fmt.Printf("dimensions: %+v\n", g)
	fmt.Println(resolution, "pixels")
}
