package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func main() {
	// request and parse the page
	resp, err := http.Get("https://www.f1fanatic.co.uk/2017/12/16/over-100-of-the-best-pictures-from-the-2017-season/")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	root, err := html.Parse(resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Printf("length: %d\n", resp.ContentLength)

	for h := range resp.Header {
		fmt.Printf("%s\n", h)
	}

	fmt.Println(resp.Header)

	articles := scrape.FindAllNested(root, nodeMatcher)

	fmt.Println("Articles:")

	for i, article := range articles {
		fmt.Printf("%2d %s (%s)\n", i, scrape.Text(article), scrape.Attr(article, "href"))
	}
}

// nodeMatcher defines a matcher for nodes
func nodeMatcher(n *html.Node) bool {
	// must check for nil values
	// if n.DataAtom == atom.A && n.Parent != nil && n.Parent.Parent != nil {
	// 	return scrape.Attr(n.Parent.Parent, "class") == "athing"
	// }

	if n.DataAtom == atom.A {
		// fmt.Println("<a> type:", n.Type)
		// fmt.Println("attr:", scrape.Attr(n, "rel"))

		for i, a := range n.Attr {
			fmt.Printf("[%3d] n: '%s'\n[%3d] k: '%s'\n[%3d] v: '%s'\n", i, a.Namespace, i, a.Key, i, a.Val)

			if a.Key == "rel" && caseInsensitiveContains(a.Val, "attachment") {
				return true
			}
		}
	}

	// fmt.Println()

	// if caseInsensitiveContains(scrape.Attr(n, "rel"), "attachment") {
	// 	fmt.Println("type:", n.Type)
	// 	// return scrape.Attr(n.Parent.Parent, "class") == "athing"
	// }

	return false
}

// caseInsensitiveContains checks if one string contains another, case-insensitively
func caseInsensitiveContains(s, substr string) bool {
	s, substr = strings.ToUpper(s), strings.ToUpper(substr)

	return strings.Contains(s, substr)
}
