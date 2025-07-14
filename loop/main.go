package main

import "fmt"

// GO only have one loop, this is for loop
func main() {
	// classic initial;condition;after for loop
	for i := 1; i <= 5; i++ {
		fmt.Println("Hello", i)
	}

	// also we can use "range" over by "do this N time"
	for l := range 3 {
		fmt.Println("Hello", l)
	}

	// infinit loop is like this
	for {
		fmt.Println("WOW, I'M INSIDE A INFINIT LOOP")
		break // we need to break to get out of this loop
	}

	// we can use continue also to the next iteration
	for n := range 10 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
