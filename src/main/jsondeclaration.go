package main

import (
	"encoding/json"
	"fmt"
	"github"
	"log"
)

// Initialization of the struct type
var movies = []Movie{
	{
		Title: "Casablanca",
		Year:  1942,
		Color: false,
		// Note that field value is same as RHS in the normal assignment
		// Note that go array is encoded in JSON as comma separated values in quotes enclosed in square bracket []
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"},
	},
	{
		Title: "Cool Hand Luke",
		Year:  1967,
		Color: true,
		// Note that field value is same as RHS in the normal assignment
		Actors: []string{"Paul Newman"},
	},
}

// Declaration of the struct array & struct on the fly without initialization
var titles []struct{ Title string }
var title struct{ Title string }

func main() {
	// It produces JSON byte slice with no whitespace
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed:%s", err)
	}
	fmt.Printf("%s\n", data)

	// json.MarshalIndent has two additional arguments
	// 1. prefix for each new line in the output byte slice
	// 2. string for each level of Indentation e.g. here it is 4 spaces wide
	// Only exported struct fields are marshaled
	data1, err1 := json.MarshalIndent(movies, "", "    ")
	if err1 != nil {
		log.Fatalf("JSON marshaling failed:%s\n", err1)
	}
	fmt.Printf("%s\n", data1)

	// Note following if condition, it is recommended only if the func return only single value like err
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("Json unmarshaling failed:%s\n", err)
	}
	fmt.Printf("%s\n", titles)

	github.SearchIssues()
}

type Movie struct {
	Title string
	// `json:...` is a json field tag i.e. string of metadata associated with field of struct at compile time
	// it consist of key value pair enclosed in `` values are often enclosed in double quotes
	Year int `json:"released"`
	// json field tags first value decides alternative JSON name for Go field name
	// second value i.e. omitempty states the behavior in case if the Go field gets its zero value
	// in this case it will not encode the Go field in json if it gets false as the value
	Color  bool `json:"color,omitempty"`
	Actors []string
}
