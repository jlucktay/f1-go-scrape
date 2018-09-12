package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"
)

func main() {
	resp, errGet := http.Get("https://jlucktay.surge.sh")
	if errGet != nil {
		log.Fatal(errGet)
	}
	defer resp.Body.Close()

	headerKeys := make([]string, 0, len(resp.Header))
	longestKey := 0

	for h := range resp.Header {
		headerKeys = append(headerKeys, h)

		if len(h) > longestKey {
			longestKey = len(h)
		}
	}

	sort.Strings(headerKeys)

	for _, key := range headerKeys {
		fmt.Printf("% -*s = %+v\n", longestKey, key, resp.Header[key])
	}
}
