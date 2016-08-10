package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		host, err := os.Hostname()
		if err != nil {
			host = "host err"
		}
		fmt.Fprintf(resp, "host: %v", host)
	})
	http.ListenAndServe(":8080", nil)
}
