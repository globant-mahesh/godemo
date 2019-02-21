package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	urls := os.Args[1:]
	for _, url := range urls {
		// net/http package is to be used for server connections
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Reading url%s: %v\n", url, err)
			os.Exit(1)
		} else {
			// struct field: resp.Body field is stream of response body
			body, err := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Reading url%s: %v\n", url, err)
			}
			fmt.Println(body)
		}
	}
}
