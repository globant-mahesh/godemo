//The naming convention for Go package is to use the name of the system directory where
// we are putting our Go source files
package main

import (
	"fmt"

	"../hello"
)

func main() {
	// The Printlnfunction is exported.
	// It starts with an upper caseletter, which means other packages are allowed to call it.
	fmt.Println("Hello 世界")
	// i := 10 + 2
	// fmt.Printf("%d", i)
	hello.Hello("World")
}
