package main

import (
	"bufio"
	"fmt"
	"os"
)

//
// All package inbuilt functions are with Intial in caps
// all inbuilt types are with initial small e.g. string (not String as in java)
func main() {
	counts := make(map[string]int)
	// Equivalent to Java: Scanner scanner = new Scanner(System.in)
	scanner := bufio.NewScanner(os.Stdin)
	// The range is explicitely taken from scanner inbuilt function Scan() function
	// this is while loop as there is only boolean condition after for keyword
	for scanner.Scan() {
		line := scanner.Text()
		counts[line]++
		if len(line) == 0 {
			break
		}
	}
	// range on the map returns key & value similar to index and value in case of array
	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%s\t%d\n", line, count)
		}
	}
}
