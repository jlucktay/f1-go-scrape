// Scrape cool F1 pictures from the web
package main

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://www.racefans.net/2017/12/16/over-100-of-the-best-pictures-from-the-2017-season/"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("User-Agent", "golang/f1-go-scrape")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
}
