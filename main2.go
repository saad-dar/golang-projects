package main

import "fmt"

//array
func todo() {
	// var arr []int

	arr := []int{1, 2, 3}

	arr2 := []string{ "hi", "my", "name" }

	arr2 = append(arr2, "is" , "saadullah",  "dar", "!")
	fmt.Println(arr, arr2)
}

func main () {
	todo()
}