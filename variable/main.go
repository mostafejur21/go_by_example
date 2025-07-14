package main

import "fmt"

func main() {
	// we can initial variable using var keyword
	var a = "Initial"
	fmt.Println(a)

	// we can declare multiple variable in single line
	var b, c int = 1, 2
	fmt.Println(b, c)

	// without initialization, go will initialize this with `zero` values
	var d int
	fmt.Println(d) // 0

	// the ':=' syntex is for declare and initialize a variable
	f := "Apple" // this is like 'var f string = "Apple"
	fmt.Println(f)
}
