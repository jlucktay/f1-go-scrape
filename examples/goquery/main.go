// From:
// https://www.progville.com/go/goquery-jquery-html-golang/
package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	doc, err := goquery.NewDocument("https://blog.golang.org")
	// Deprecated: Use the net/http standard library package to make the request and validate the response before calling goquery.NewDocumentFromReader with the response's body.

	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".article").Each(func(i int, s *goquery.Selection) {
		title := s.Find("h3").Text()
		link, _ := s.Find("h3 a").Attr("href")
		fmt.Printf("%d) %s - %s\n", i+1, title, link)
	})
}
