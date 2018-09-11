package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
)

func main() {
	req, errReq := http.NewRequest("GET", "http://httpbin.org/user-agent", nil)
	if errReq != nil {
		log.Fatal(errReq)
	}

	req.Header.Set("User-Agent", "Golang_Spider_Bot/"+runtime.Version())

	resp, errGet := http.DefaultClient.Do(req)
	if errGet != nil {
		log.Fatal(errGet)
	}
	defer resp.Body.Close()

	body, errRead := ioutil.ReadAll(resp.Body)
	if errRead != nil {
		log.Fatal(errRead)
	}

	fmt.Println(string(body))
}
