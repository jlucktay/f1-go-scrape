package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Set your email here to include in the User-Agent string.
var email = "youremail@gmail.com"
var urls = []string{
	"https://buzzfeed.com",
	"https://digg.com",
	"https://en.wikipedia.org",
	"https://news.ycombinator.com",
	"https://reddit.com",
	"https://techcrunch.com",
}

func respGen(urls ...string) <-chan *http.Response {
	var wg sync.WaitGroup
	out := make(chan *http.Response)
	wg.Add(len(urls))

	for _, url := range urls {
		go func(url string) {
			req, errReq := http.NewRequest("GET", url, nil)
			if errReq != nil {
				log.Fatal(errReq)
			}

			req.Header.Set("User-Agent", "testBot("+email+")")

			resp, errGet := http.DefaultClient.Do(req)
			if errGet != nil {
				log.Fatal(errGet)
			}

			out <- resp
			wg.Done()
		}(url)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func rootGen(in <-chan *http.Response) <-chan *html.Node {
	var wg sync.WaitGroup
	out := make(chan *html.Node)

	for resp := range in {
		wg.Add(1)

		go func(resp *http.Response) {
			root, err := html.Parse(resp.Body)

			if err != nil {
				log.Fatal(err)
			}

			out <- root
			wg.Done()
		}(resp)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func titleGen(in <-chan *html.Node) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)

	for root := range in {
		wg.Add(1)

		go func(root *html.Node) {
			title, ok := scrape.Find(root, scrape.ByTag(atom.Title))

			if ok {
				out <- scrape.Text(title)
			}

			wg.Done()
		}(root)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	// Set up the pipeline to consume back-to-back output ending with the final
	// stage to print the title of each web page in the main go routine.
	for title := range titleGen(rootGen(respGen(urls...))) {
		fmt.Println(title)
	}
}
