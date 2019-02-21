package main

import (
	"fmt"
	"os"
)

func main() {

	// Explicit initialization
	str, sep := "", ""
	// use of blank identifier in for loop
	for _, arg := range os.Args[1:] {
		str += sep + arg
		sep = " "
	}
	fmt.Println(str)
}
