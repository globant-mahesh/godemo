package main

import (
	"fmt"
)

// func main() {
// 	i := 10
// 	rune := 'A'
// 	// Note initialisation of struct is comma separated json like path
// 	b := book{
// 		"Mybook",
// 		100,
// 	}
// 	fmt.Printf("%T\n", i)
// 	fmt.Println(reflect.TypeOf(i))
// 	fmt.Println(b)
// 	fmt.Println(rune)
// 	fmt.Println(reflect.TypeOf(rune))
// 	fmt.Println(runtime.NumCPU())
// }

type book struct {
	bookName  string
	pageCount int
}

type reader interface {
	getName()
}

// Inheritance
func (b book) getName() {
	fmt.Println(b.bookName)
}
