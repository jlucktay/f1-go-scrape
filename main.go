// Scrape cool F1 pictures from the web
package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	targetURL, _ := url.Parse("https://www.racefans.net/2017/12/16/over-100-of-the-best-pictures-from-the-2017-season/")

	res, err := http.DefaultClient.Do(newRequest(*targetURL))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// doc.Find("img").Each(func(i int, s *goquery.Selection) {
	// 	src := s.AttrOr("src", "")
	// 	alt := s.AttrOr("alt", "")
	// 	fmt.Printf("'%s' <%s>\n", alt, src)
	// })

	// .Html()
	// .Length()
	// .Text()

	counter := 0

	for img := range genImageUrlsFromDoc(doc) {
		fmt.Printf("[%d] img: '%s'\n", counter, img.String())
		counter++
	}
}

func genImageUrlsFromDoc(input *goquery.Document) chan url.URL {
	output := make(chan url.URL)

	go func() {
		input.Find("img").Each(func(i int, s *goquery.Selection) {
			src := s.AttrOr("src", "")
			if u, _ := url.Parse(src); u.IsAbs() && u.Scheme != "data" {
				output <- *u
			}
		})

		close(output)
	}()

	return output
}

/*
<img
	src="https://www.f1fanatic.co.uk/wp-content/uploads/2017/03/XPB_865994_HiRes.jpg"
	alt="Daniel Ricciardo, Red Bull, Albert Park, 2017"

	class="size-full wp-image-338264"
	srcset="
		https://www.racefans.net/wp-content/uploads/2017/03/XPB_865994_HiRes.jpg 1920w
		,
		https://www.racefans.net/wp-content/uploads/2017/03/XPB_865994_HiRes-470x313.jpg 470w
		,
		https://www.racefans.net/wp-content/uploads/2017/03/XPB_865994_HiRes-768x512.jpg 768w
		,
		https://www.racefans.net/wp-content/uploads/2017/03/XPB_865994_HiRes-886x591.jpg 886w
	"
	sizes="(max-width: 1920px) 100vw, 1920px"
>
*/

func genFilterLargeImageUrls(input chan url.URL) chan url.URL {
	output := make(chan url.URL)

	go func() {
		for i := range input {

			/*
				resp, err := http.Get("http://i.imgur.com/Peq1U1u.jpg")
				if err != nil {
					log.Fatal(err)
				}
				defer resp.Body.Close()

				m, _, err := image.Decode(resp.Body)
				if err != nil {
					log.Fatal(err)
				}
				g := m.Bounds()

				// Get height and width
				height := g.Dy()
				width := g.Dx()

				// The resolution is height x width
				resolution := height * width

				// Print results
				fmt.Println(resolution, "pixels")
			*/

			/*
			   res, err := http.DefaultClient.Do(newRequest(url))
			   if err != nil {
			   	log.Fatal(err)
			   }
			   defer res.Body.Close()

			   if res.StatusCode != 200 {
			   	log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
			   }
			*/

			output <- i

		}

		close(output)
	}()

	return output
}

func newRequest(u url.URL) *http.Request {
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("User-Agent", "golang/f1-go-scrape")

	return req
}
