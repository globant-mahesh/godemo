package main
//
//import (
//	"fmt"
//	"os"
//	"strconv"
//	"types"
//)
//
//func main() {
//	var i uint64
//	var f float32
//	var f2 float64
//	fmt.Println("Hello, World")
//	fmt.Println("Hello There")
//	fmt.Println(i, ",", f, ",", f2)
//
//	var a int32 = 10
//	if a%2 == 0 {
//		fmt.Printf("%d is a Even Number\n", a)
//	} else {
//		fmt.Println("Odd Number")
//	}
//
//	var b int64 = 321
//	switch b {
//	case 1:
//		fmt.Printf("Value is %d\n", 1)
//		fallthrough
//	case 321:
//		fmt.Printf("Values is %d\n", 321)
//		fallthrough
//	default:
//		fmt.Printf("Print Again\n")
//	}
//
//	for a := 0; a < 3; a++ {
//		for b := 4; b > 0; b-- {
//			fmt.Printf("%d\t%d\n", a, b)
//		}
//	}
//
//	num := []int{1, 4, 9}
//	for _, value := range num {
//		fmt.Println(value)
//	}
//
//	myMap := map[string]string{"name": "mahesh", "middle": "a", "last": "bhosale"}
//
//	for k, v := range myMap {
//		fmt.Println(k, v)
//	}
//
//	const MYINT int = 10
//
//	fmt.Println(MYINT)
//
//	var j float64 = 87.23112
//	var k int64 = 23
//	var str string = "145.789"
//	strnew, _ := strconv.ParseFloat(str, 64)
//
//	fmt.Println(int32(strnew), int32(j), float32(k))
//
//	e := Employee{"Mahesh", "Bhosale"}
//
//	firstName, lastName, middleName := e.getDetails("A")
//
//	fmt.Println(firstName, lastName, middleName)
//
//	addition := addAll(10, 78, 23, 67)
//	fmt.Println(addition)
//
//	fmt.Println(factorial(5))
//
//	square := func() int {
//		square := MYINT * MYINT
//		return square
//	}
//	fmt.Println(square)
//
//	fmt.Println(additionMethod(12, 78))
//
//	fmt.Println(initializeTwoDArray())
//
//	slicea, sliceb := initializeSlice()
//
//	printSlice(slicea)
//	printSlice(sliceb)
//	printCommandLineArgs()
//	createStructAndDisplay()
//
//	e := employee{p, 123}
//	p := person{"Mahesh", "Bhosale"}
//	p.getDeails()
//	e.getDetails()
//	//	ar := [...]int{1, 2, 3}
//	//	Slice variable has predefined capacity of underlying array here it is 3
//	// Also its length is 3
//	//	s := ar[:3]
//	//	s := []int{}
//	s := make([]int, 0, 1)
//	//	We can change the capacity of the slice through function
//	// slice is passed as call by reference to function so it can change its value within function
//	// Also we can get the modified slice in the same slice varibale as the returned value changing its original capacity
//	s = appendToSlice(s, 4)
//	fmt.Println(s)
//}
//
//type Employee struct {
//	firstName string
//	lastName  string
//}
//
////	Function of struct
//func (e Employee) getDetails(middleName string) (string, string, string) {
//	return e.firstName, e.lastName, middleName
//}
//
////	function as method i.e. receiver of arguments
//func addAll(args ...int) int {
//	addition := 0
//	for e, i := range args {
//		addition += i
//		fmt.Println(e, i)
//	}
//	return addition
//}
//
//func factorial(num int) int {
//	if num == 0 || num == 1 {
//		return 1
//	} else {
//		return num * factorial(num-1)
//	}
//
//}
//
////closure
//func additionMethod(a int, b int) int {
//	incrementValue := 1
//	//closure - anonymous function within another fucntion capable of referring variables declared outside its body
//	closureIncrement := func() int {
//		return a + incrementValue
//	}
//	return (b + closureIncrement())
//}
//
//func initializeTwoDArray() [2][3]int {
//	var a [2][3]int
//	a = [2][3]int{
//		//default value zero
//		{1, 8},
//		{9, 34, 56},
//	}
//	return a
//}
//
////create slice
////slice is like an array declared without setting the size having access to underlying true array
//func initializeSlice() ([]int, []int) {
//	a := []int{
//		12, 34,
//	}
//	b := make([]int, 5, 10)
//	return a, b
//}
//
//func printSlice(slice []int) {
//	fmt.Printf("Length of slice %d & capacity of slice = %d slice=%v\n", len(slice), cap(slice), slice)
//}
//
//func printCommandLineArgs() {
//	fmt.Printf("%s %s %s\n", os.Args[1], os.Args[2], os.Args[3])
//	for _, str := range os.Args[1:] {
//		fmt.Println(str)
//		str = "mystr"
//		fmt.Println(str)
//	}
//}
//
//func createStructAndDisplay() {
//	type person struct {
//		firstName string
//		lastName  string
//		id        int
//	}
//	x := person{firstName: "Mahesh", lastName: "Bhosale", id: 123}
//	fmt.Println(x)
//}
//
//type person struct {
//	firstName string
//	lastName  string
//}
//type employee struct {
//	person
//	id int
//}
//
//func (p person) getDetails() {
//	fmt.Println(p.firstName, " ", p.lastName)
//}
//func (e employee) getDetails() {
//	fmt.Println(e.person, " ", e.id)
//}
//
//func appendToSlice(x []int, y int) []int {
//	var z []int
//	len := len(x) + 1
//	if len <= cap(x) {
//		z = x[:len]
//	} else {
//		//		Following condition is considering the capacity of the empty slice i.e. capacity = 0
//		//		Capacity of the slice is always >= length of the slice so oloqing if condition becomes redundant
//		//      as it is already handled by the enclosing else block
////		if len > cap(x) {
//			c := 2 * len
//			z = make([]int, len, c)
////		}
//		copy(z, x)
//	}
//	return z
//}
