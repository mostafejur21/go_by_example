package main

import "fmt"

// here, this intSeq function is returning a anonymous function that will form
// a closure with the variable i
func intSeq() func() int {
	i := 0
	return func() int {
		fmt.Println(&i)
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
	fmt.Println(newInts())
}
