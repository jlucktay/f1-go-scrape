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
	resp, errGet := http.Get("http://i.imgur.com/Peq1U1u.jpg")
	if errGet != nil {
		log.Fatal(errGet)
	}
	defer resp.Body.Close()

	m, f, errDecode := image.Decode(resp.Body)
	if errDecode != nil {
		log.Fatal(errDecode)
	}

	fmt.Println("Format:", f)
	g := m.Bounds()

	// The resolution is height x width
	resolution := g.Dy() * g.Dx()

	// Print results
	fmt.Printf("Dimensions: %+v\n", g)
	fmt.Println(resolution, "pixels")
}
