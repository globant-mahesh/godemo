package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
 * defer panic recover
 * https://blog.golang.org/defer-panic-and-recover
 */
// func main() {
// 	// panic is equivalent to throw new Exception() in java
// 	// panic("First panic.")
// 	checkIfEligibeforVote()
// }

func checkIfEligibeforVote() {
	// defer functions are called when enclosing function returns (a value if has a return type)
	// defer function calls are pushed on to the stack so follow LIFO order
	// recover regains control of panicking goroutine & only useful if it is used in the first deferred
	// function inside the goroutine(enclosing function)
	// Anonymous functions are used in the deferred function
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	fmt.Println("Enter the age of the candidate")
	scanner := bufio.NewScanner(os.Stdin)
	// for with only condition mentioned to make sure that there is no break in the input data
	for scanner.Scan() {
		str := scanner.Text()
		age, _ := strconv.ParseInt(str, 0, 64)
		//deferred function 2 will be on top of deferred functions stack in the enclosing function
		// i.e. goroutine
		defer fmt.Printf("Entered age was: %d\n", age)
		if age > 18 {
			fmt.Println("Eligible for voting.")
			break
		} else if age > 0 && age < 18 {
			fmt.Println("Not Eligible for voting.")
			break
		} else if age < 0 {
			fmt.Println("You have mistakened.")
			panic("Invalid age.")
		}
	}
}
