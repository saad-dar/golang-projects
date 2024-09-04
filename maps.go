package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]

	fmt.Println("v1: ", v1)

	v2 := m["k3"]

	fmt.Println("v2: ", v2)

	fmt.Println("len: ", len(m))

	delete(m , "k2")

	fmt.Println("maps: ", m)

	clear(m)

	fmt.Println("map:", m)

	// The optional second return value when getting a value from a map indicates if 
	// the key was present in the map. This can be used to disambiguate between missing 
	// keys and keys with zero values like 0 or "". Here we didn’t need the value itself,
	//  so we ignored it with the blank identifier _
	_, prs := m["k2"] 
	fmt.Println("prs", prs)

	n := map[string]int{ "foo": 1, "bar": 2}

	// Note that maps appear in the form map[k:v k:v] when printed with fmt.Println.
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}

	// The maps package contains a number of useful utility functions for maps.
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}