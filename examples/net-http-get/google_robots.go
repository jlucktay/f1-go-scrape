package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, errGet := http.Get("https://www.google.com/robots.txt")
	if errGet != nil {
		log.Fatal(errGet)
	}
	defer res.Body.Close()

	robots, errRead := ioutil.ReadAll(res.Body)
	if errRead != nil {
		log.Fatal(errRead)
	}

	fmt.Println(string(robots))
}
