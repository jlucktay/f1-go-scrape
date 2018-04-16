package main

import (
	"fmt"
	"net/http"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func main() {
	// request and parse the front page
	resp, err := http.Get("https://news.ycombinator.com/")

	if err != nil {
		panic(err)
	}

	root, err := html.Parse(resp.Body)

	if err != nil {
		panic(err)
	}

	// grab all articles and print them
	articles := scrape.FindAll(root, nodeMatcher)

	for i, article := range articles {
		fmt.Printf("%2d %s (%s)\n", i, scrape.Text(article), scrape.Attr(article, "href"))
	}
}

// <a href="https://www.atlasobscura.com/places/encryption-lava-lamps" class="storylink">The randomness of this wall of lava lamps helps encrypt up to 10% of internet</a>

// article
// <*golang.org/x/net/html.Node>

// article.Attr
// <[]golang.org/x/net/html.Attribute> (length: 2, cap: 2)

// article.Attr[0]
// <golang.org/x/net/html.Attribute>
// Namespace:""
// Key:"href"
// Val:"https://www.atlasobscura.com/places/encryption-lava-lamps"

// article.Attr[1]
// <golang.org/x/net/html.Attribute>
// Namespace:""
// Key:"class"
// Val:"storylink"

// define a matcher
func nodeMatcher(n *html.Node) bool {
	// must check for nil values
	if n.DataAtom == atom.A && n.Parent != nil && n.Parent.Parent != nil {
		return scrape.Attr(n.Parent.Parent, "class") == "athing"
	}

	return false
}
