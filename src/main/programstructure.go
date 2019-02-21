package main

import (
	"fmt"
	. "fmt"
)

const boilingF = 312.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
	// Output:
	// boiling point = 212°F or 100°C
	// Direct definition of array
	q := [...]int{1, 2, 3}
	for _, v := range q {
		fmt.Println(v)
	}

	arra := [...]int{1, 2}
	arrb := [2]int{1, 2}
	arrc := [2]int{2, 3}
	// arrd := [3]int{1, 2}
	Println(arra == arrb, arra == arrc)
	// both array type & array values will be compared for following condition
	// Println(arra == arrd)
}
