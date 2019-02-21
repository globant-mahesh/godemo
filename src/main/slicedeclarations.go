package main

import "fmt"

func main() {

	// Type 1 general array like declaration
	slice := []int{10, 23, 34}
	fmt.Println(len(slice), cap(slice), slice)

	// Type 2 make function set inidividual element using indices/for loop
	slice1 := make([]int, 4, 7)
	slice1[0] = 1
	slice1[1] = 2
	slice1[2] = 3
	slice1[3] = 4
	// compilation error
	// slice1[4] = 5
	fmt.Println(slice1)

	// Type 2 for initialization use make function & use copy inbuilt function
	// which will hold the value for the set length in make not beyond it if set
	// to zero nothing will be copied
	slice2 := make([]int, 3, 8)
	copy(slice2, slice1)
	fmt.Println(slice2)
	//Best way is by declaring new slice with the length of existing slice as
	// capacity & length as zero & use append func in place of copy
	slice2mod := make([]int, 0, len(slice1))
	for _, val := range slice1 {
		slice2mod = append(slice2mod, val)
	}
	fmt.Println(slice2mod)

	// Type 3: slice notation on array or existing slice
	arr1 := [4]int{12, 12, 43, 87}
	slice3 := arr1[:3]
	fmt.Println(len(slice3), cap(slice3), slice3)
	slice4 := slice3[:1]
	fmt.Println(len(slice4), cap(slice4), slice4)
}
