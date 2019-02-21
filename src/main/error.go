package main

import (
	"errors"
	"math"
)

func sqrt(i float64) (float64, error) {
	if i < 0 {
		// Use errors package to create errors (like throwing an exception in java)
		return 0, errors.New("Square root of the negative number")
	}
	return math.Sqrt(i), nil
}

// func main() {
// 	i, err := sqrt(-64)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(i)
// 	}

// 	j, err := sqrt(64)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(j)
// 	}
// }
