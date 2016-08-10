package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// connect app1
	resp, err := http.Get("http://app1:8080/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", b)
}
