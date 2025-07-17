package main

import "fmt"

// here the sum function that will take an arbirtary numbers of ints as arguments
// means, it can takes any amount of int as an arguments. no limit, this type of
// function is called Variadic Function
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 3)
	sum(1, 2, 3)
	sum(1, 2, 3, 4, 5)
	sum(1, 2, 3, 5, 8, 9)
	nums := []int{1, 3, 4, 5, 6, 7, 8, 9}
	sum(nums...)
}
