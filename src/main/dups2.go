package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	files := os.Args[1:]
	counts := make(map[string]int)
	for _, filename := range files {
		// io/ioutill is the package for read/write from file or stream
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\t", err)
		} else {
			lines := strings.Split(string(data), "\n")
			for _, line := range lines {
				counts[line]++
			}
			for line, count := range counts {
				if count > 1 {
					fmt.Printf("%s\t%d\n", line, count)
				}
			}
		}
	}
}
