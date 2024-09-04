// Accepting a variable number of arguments is another nice feature of Go functions

package main

import "fmt"

func sum(nums ...int) int {
	fmt.Println(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println("total", total)
	return total
}

func main () {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	total := sum(nums...)
	fmt.Println(total)
}