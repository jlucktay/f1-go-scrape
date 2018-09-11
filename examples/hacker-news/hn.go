package main

import (
	"fmt"
	"net/http"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func main() {
	// Request and parse the front page
	resp, getErr := http.Get("https://news.ycombinator.com/")
	if getErr != nil {
		panic(getErr)
	}

	root, parseErr := html.Parse(resp.Body)
	if parseErr != nil {
		panic(parseErr)
	}

	// Grab all articles and print them
	articles := scrape.FindAll(root, nodeMatcher)

	for i, article := range articles {
		fmt.Printf("%2d) %s (%s)\n", i, scrape.Text(article), scrape.Attr(article, "href"))
	}
}

// Define a matcher
func nodeMatcher(n *html.Node) bool {
	// Must check for nil values
	if n.DataAtom == atom.A && n.Parent != nil && n.Parent.Parent != nil {
		return scrape.Attr(n.Parent.Parent, "class") == "athing"
	}

	return false
}
