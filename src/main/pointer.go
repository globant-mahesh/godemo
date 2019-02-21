package main

import (
	"flag"
	"fmt"
	"strings"
)

// First argument of the following flag fucntions serve as the command line option(flag) names
// Return value is pointer to command line option vaue
var n = flag.Bool("n", false, "Omit Trailing newline")
var sep = flag.String("s", " ", "Separater")

func main() {
	//	type var
	i := 10
	changevalue(&i)
	fmt.Println(i)

	//	Address var
	j := new(int)
	*j = 10
	fmt.Println(*j, j)
	changevalue(j)
	fmt.Println(*j, j)

	// function with return value as pointer to local variable returns
	// different value on each call
	fmt.Println(f() == f())

	k := 34
	k = increment(&k)
	fmt.Println(k)
	fmt.Println(increment(&k))

	echo4()

	// Use of new function for unnamed variable creation
	a := new(int)
	b := new(int)
	fmt.Println(a, b, *a, *b, a == b, *a, *b, *a == *b)
}

func echo4() {
	// echo command line arguments using flag package using flag.Args()
	// It uses pointers declared in the global variables as first argument
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

func changevalue(i *int) {
	*i = 20
}

func f() *int {
	i := 2
	return &i
}

func increment(count *int) int {
	*count++
	return *count
}
