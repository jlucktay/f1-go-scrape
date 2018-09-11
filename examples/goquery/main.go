// From:
// https://www.progville.com/go/goquery-jquery-html-golang/
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	res, getErr := http.DefaultClient.Get("https://blog.golang.org")
	if getErr != nil {
		log.Fatal(getErr)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Status code error: %d\n%s\n", res.StatusCode, res.Status)
	}

	doc, readErr := goquery.NewDocumentFromReader(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	doc.Find(".article").Each(func(i int, s *goquery.Selection) {
		title := s.Find("h3").Text()
		link, _ := s.Find("h3 a").Attr("href")
		fmt.Printf("%d) %s - %s\n", i+1, title, link)
	})
}
