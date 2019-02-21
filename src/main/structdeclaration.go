package main

import (
	"fmt"
)

// var leftmost *node = new(node)
var w Wheel

func main() {
	// Declaration of array
	values := [5]int{23, 12, 13, 67, 34}
	// Declaration of pointer to struct
	tr := new(tree)
	// Note the array type is converted into slice
	tr.insertNode(values[:])
	fmt.Println(tr.start.value, tr.tail.value)
	// appendValues(values, &tr)
	// **declaration of empty struct equivalent to empty array e.g. []int{}
	// **declaration of struct initialising each field value i.e. struct literal
	// type Point struct{ X, Y int }
	// Defining struct literal using only values. Order must be as per the struct declaration.
	p := Point{1, 2}
	// Defining struct literal using field names explicitely. Fields not declared are initialsed to their zero value.
	p1 := Point{Y: 11}
	// Shorthand for creation initialisaztion of struct & exctracting address in it's pointer
	// Very useful for creating & passing the reference on the fly in function call as an argument
	p2 := &Point{Y: 23}
	// Equivalent to
	// new is inbuilt function so starts with small letter 'n'
	// can not be used for passing as an argument expression
	p3 := new(Point)
	*p3 = Point{43, 56}
	p4 := &Point{23, 12}
	fmt.Println(p, p1, *p2, p3, *p4)
	// If ALL the fields of struct type are comparable then struct is comparable using == & !=
	fmt.Println(p.X == p1.X, p.Y == p1.Y, p == p1, p == p)

	recordResidents()
	// Equivalent of w.Circle.Point.X
	w.X = 12
	// Equivalent of w.Circle.Point.Y
	w.Y = 23
	w.Radius = 12.1
	w.Spokes = 24
	fmt.Println(w)

	// JSON like initialization must include all field values
	// Field name of anonymous field here becomes its type name from RHS of ":"
	w = Wheel{
		Circle: Circle{
			Point:  Point{X: 31, Y: 22},
			Radius: 67.1,
		},
		Spokes: 24,
	}
	fmt.Println(w)

	// Non JSON like initialization includes only structs
	// Note inner non struct fields refer the non JSON single value syntax
	w = Wheel{Circle{Point{12, 38}, 23.1}, 24}
	// adverb "#" adds syntax of golang in representation of the %v verb value
	fmt.Printf("%#v\n", w)
}

// Note that there is no comma separation on the struct type field declarations
type Point struct {
	// X & Y can be exported
	X int
	Y int
}

// Use of anonymous field to avoid chain of variables in dot expression due to
// structs embedded in another struct
type Circle struct {
	// Only type(treat the RHS of field i.e. type NOT name for following literal) is declared for anonymmous field
	// Anonymous field has implicit name though same as type name
	// e.g. implicit name of anonymous field Point is Point
	// You can not have two anonymous fields of same type as it would lead to the conflict on getting implicit field name
	// Anonymous fields are by default exported iff their type name starts with capital letter.
	Point
	Radius float32
}

type Wheel struct {
	// Anonyomus field i.e. only type is specified
	Circle
	Spokes int
}

// ONLY comparable struct types can be used as key type of the map
type address struct {
	hostName   string
	flatNumber int
}

// Function within function is not allowed except closure/anonymous function
func recordResidents() {
	m := make(map[address]int)
	m[address{"Mahesh", 101}]++
	fmt.Println(m)
}

// Order of fields of struct aggregate type decides its unique identity
// It can contain field pointer to it's own type
// It can NOT contain it's own type as field type
// The field starting with capital etter name are exported(public) others are not (private)
type node struct {
	value      int
	next, prev *node
}

type tree struct {
	start, tail *node
}

// Use struct pointer to make the function callable from struct pointer reference
// Call by reference i.e. passing pointer to the struct as an argument adds to efficiency in larger structs
func (t *tree) insertNode(values []int) {
	for _, v := range values {
		newNode := new(node)
		newNode.value = v
		if t.start == nil {
			t.start = newNode
			t.tail = newNode
		} else {
			currentNode := t.start
			for currentNode.next != nil {
				currentNode = currentNode.next
			}
			// Always initialise pointers with pointer object before changing the pointer
			newNode.prev = currentNode
			currentNode.next = newNode
			t.tail = newNode
		}
	}
}

// func sort(t *node, values []int) {
// 	for _, v := range values {
// 		t = add(t, v)
// 	}
// 	leftmost := t
// 	for leftmost.left != nil {
// 		leftmost = leftmost.left
// 	}
// 	t = leftmost
// 	// leftmost = leftmost.right
// 	// for leftmost.right != nil {
// 	// 	fmt.Println(leftmost)
// 	// 	leftmost = leftmost.right
// 	// }
// 	// Pass empty slice to change underlying array saving memory
// 	appendValues(values[:0], t)
// }

// func appendValues(values []int, t *node) []int {
// 	if t != nil {
// 		values = appendValues(values, t.left)
// 		values = append(values, t.value)
// 		values = appendValues(values, t.right)
// 	}
// 	return values
// }

// func add(t *node, v int) *node {
// 	if t == nil {
// 		// Create a pointer to node
// 		t := new(node)
// 		// assignment of the field with = operator not with :=
// 		t.value = v
// 		return t
// 	} else if v < t.value {
// 		// Recursion can be used in place of following 2 steps
// 		// t.left = new(node)
// 		// t.left.value = v
// 		// Pass pointer to nil in add to trigger new node struct creation
// 		if t.left == nil {
// 			t.left = add(t.left, v)
// 			// t.left.right = t
// 		} else {
// 			// Create a pointer to node
// 			t1 := new(node)
// 			// assignment of the field with = operator not with :=
// 			t1.value = v
// 			if t.left.value < t1.value {
// 				t1.left = t.left
// 				t.left.right = t1
// 				t.left = t1
// 				t1.right = t
// 			}
// 		}
// 	} else {
// 		// Recursion can be used in place of following 2 steps
// 		// t.right = new(node)
// 		// t.right.value = v
// 		if t.right == nil {
// 			t.right = add(t.right, v)
// 			// t.right.left = t
// 		} else {
// 			// Create a pointer to node
// 			t1 := new(node)
// 			// assignment of the field with = operator not with :=
// 			t1.value = v
// 			if t.right.value > t1.value {
// 				t1.right = t.right
// 				t.right.left = t1
// 				t.right = t1
// 				t1.left = t
// 			}
// 		}
// 		// t.right.left = t
// 	}
// 	// on returning for initial recursion i.e. completing the formation of node what to return?
// 	if t.left != nil {
// 		return t.left
// 	} else {
// 		return t
// 	}
// 	// return t.left
// }
