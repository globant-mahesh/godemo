package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%d\n", add(12, 13, 15))
	fmt.Println(sub(25, 17))
	// fmt.Println(first(12, 22))
	fmt.Println(zero(12, 98))

}

func add(x, y int) int { return x + y }

// Mixing of named and unnamed result/input parameters is not  allowed
// In named result parameters bare return is used which asumes that return statement returns all
// result parameter in the declared order
func sub(x int, y int) (z int) {
	z = x - y
	return
}

// func first(x int, _ int) {
// 	return x
// }

func zero(int, int) int {
	return 0
}
