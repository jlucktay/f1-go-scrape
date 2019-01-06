package main

import (
	"net/url"
	"testing"

	"github.com/jlucktay/f1-go-scrape/examples/triplebyte/crawler"
)

func TestChallenge3(t *testing.T) {
	// Usually crawl with 5 threads, but we boost it here to fail faster
	c := crawler.Crawler{
		Threads: 12,
		Log:     crawler.Verbose(),
	}
	graph, err := c.Crawl("http://triplebyte.github.io/web-crawler-test-site/test3/", "")
	if err != nil {
		t.Fatal("Broken test, can't run crawl")
	}

	u, err := url.Parse("http://blah.com:7091")
	if err != nil {
		t.Fatal("Broken test, can't parse URL")
	}

	if _, ok := graph.Nodes[*u]; !ok {
		t.Errorf("Not found: %v", u)
	}
}
