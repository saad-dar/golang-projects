package main

import (
	"fmt"
	"strings"
)


func main() {
	// var m1 string

	// m1 = "my name is saadullah dar"

	// daclarazation and initialization at the same time

	m1 := "my name is saadullah dar"

	/**
	// strings are mutable in go
	m1 = "hello there"
	*/
	// string functions 

	m2 := "saadullah"

	fmt.Println(strings.Contains(m1, m2))
	fmt.Println(strings.Replace(m1, "saad", "zabeeh", 1))
	fmt.Println(strings.Split(m1, " "))
	fmt.Println(m1 + " " + m2)
}
// func main() {

// 	// var (
// 	// 	m1 int32
// 	// 	m2 int64
// 	// )

// 	// direct initialization

// 	m1 := 3
// 	m2 := 4
// 	fmt.Println("hello world!", m1 + m2)
// }