// Command line arguments can be accessed using os pacakge Args slice
// It makes the package run in the os independent manner
// Slices can be accessed like indexed array and with [:] syntax too

package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	j, k := 0, 0

	// Only Loop statement in Go is for loop as follows
	for i := 1; i < len(os.Args); i++ {
		// For the very first iteration sep is initialized to empty string & not the space.
		s += sep + " " + os.Args[i]
		// initialization of the variable to space string
		sep = " "
		// Increment & decrement operations are statements here not the expressions
		// hence, in Go, it is illegal to use ++j in place of
		j++
		// it is illgal to use k = j++ in place of
		k = j
	}
	fmt.Println(s)
	fmt.Println("Number of arguments: ", k)
}
