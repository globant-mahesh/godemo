package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {

	// General declaration with initialization
	map1 := map[string]int{"abc": 10, "xyz": 20, "def": 30}
	fmt.Println(map1)

	// Using make function plus indexed initialization
	map2 := make(map[string]int)
	map2["wer"] = 2
	map2["yus"] = 8
	map2["rty"] = 10
	//Retrival of non existing key value returns default 0 value but doesn't
	//create map entry unless there is assignment involved
	fmt.Println(map2, map2["non_existing_key"])
	fmt.Println(map2)
	map2["non_existing_key"] += 1
	map2["non_existing_key2"]++
	fmt.Println(map2)

	// Iterating over map entries for range syntax similar to arrays/slices
	for key, value := range map2 {
		fmt.Printf("%s %d\n", key, value)
	}

	// Iterating over only map keys is also possible with for range
	slice := []string{}
	for key := range map2 {
		// append function has return type to take advantage of reutilization of
		// underlying array
		slice = append(slice, key)
	}
	fmt.Println(slice)

	//Iterating over sorted(ordered) keys
	sort.Strings(slice)
	for _, sortedKey := range slice {
		fmt.Println(map2[sortedKey])
	}

	// This is declaration with initialization
	map3 := make(map[string]int)
	fmt.Println(len(map3))
	map3["asdas"] = 3
	fmt.Println(map3)

	// This is only declaration that results in nil map
	var map4 map[string]int
	fmt.Println(len(map4))
	// nil map assignment results in panic
	// map4["reefd"] = 4
	// use make function to initialize it default length i.e. 0 length
	map4 = make(map[string]int)
	map4["reefd"] = 4
	fmt.Println(map4)

	// Check if key really exist in the map
	if val, ok := map4["reefd"]; ok {
		fmt.Println(val, map4["not_existing_key"])
	}

	// Note the general declaration initialization
	map5 := map[string]int{"reefd": 4}
	fmt.Println(compareMap(map4, map5))
	fmt.Println(compareMap(map4, map3))
}

// Comparing two maps is not possible with == or any inbuilt method except with nil map
// It can be done as follows
func compareMap(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, v := range x {
		if val, ok := y[k]; !ok || val != v {
			return false
		}
	}
	return true
}

// Use of map as a set
func dedup() {
	m1 := make(map[string]bool)
	scanner := bufio.NewScanner(os.Stdin)
	// scanner does't require range on scan method
	for scanner.Scan() {
		// Get string value of last scan
		line := scanner.Text()
		if !m1[line] {
			fmt.Println(line)
		}
	}
	// Error handing in IO operation
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(code)
	}
}

// Use non comparable types in map as Keys like slice
func getcomparablekey(slice []int) string {
  return fmt.Sprintf("%q", slice)
}
var mapofstring = map[string]int
func add(slice []int) {
  mapofstring[getcomparablekey(slice)]++
}
func count(slice []int) int {
  return mapofstring[getcomparablekey(slice)] 
}
