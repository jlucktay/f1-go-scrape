package main

import (
	"fmt"
	"net/http"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func main() {
	// request and parse the page
	resp, err := http.Get("https://jlucktay.surge.sh")
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	root, err := html.Parse(resp.Body)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Printf("length: %d\n", resp.ContentLength)

	for h, j := range resp.Header {
		fmt.Printf("'%s' = '%s'\n", h, j)
	}

	fmt.Println(resp.Header)

	// grab all articles and print them
	articles := scrape.FindAllNested(root, nodeMatcher)

	fmt.Println("Articles:")

	for i, article := range articles {
		fmt.Printf("%2d %s (%s)\n", i, scrape.Text(article), scrape.Attr(article, "href"))
	}
}

func nodeMatcher(n *html.Node) bool {
	// must check for nil values
	if n.DataAtom == atom.A && n.Parent != nil && n.Parent.Parent != nil {
		return scrape.Attr(n.Parent.Parent, "class") == "athing"
	}

	return false
}
