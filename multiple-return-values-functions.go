
// Go has built-in support for multiple return values. 
// This feature is used often in idiomatic Go, 
// for example to return both result and error values from a function.

package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()

	fmt.Println(a, b)

	_, c := vals() // we can ingore values using underscore(_)
	fmt.Println(c)

	d, _ := vals()
	fmt.Println(d)
}