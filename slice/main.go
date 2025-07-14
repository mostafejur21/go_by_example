package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string // this will make a empty slice with length of zero (0)
	fmt.Println("uninit", s, s == nil, len(s) == 0)

	// we can create slice with non-empty length using builtin make function
	s = make([]string, 3) // this will create a slice length of 3
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// we can set and get just like array
	s[0] = "A"
	s[1] = "B"
	s[2] = "C"
	fmt.Println("Set:", s)
	fmt.Println("Get:", s[2])

	// we can append slice with new values, which may return a value
	s = append(s, "D")
	s = append(s, "E", "F")
	fmt.Println("apd", s)

	// we can copy a slice into another slice
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy", c)

	// we can slice a Slice with this
	l := s[2:5] // the last one (upper bound) is excluding
	fmt.Println("slice[2:5]:", l)

	f := s[:5] // from the start to the end (exclude 5)
	fmt.Println("slice[:5]:", f)

	g := s[1:] // from the 1 to the end (here 1 is include)
	fmt.Println("slice[1:]:", g)


	// we can declare and initialize within a single line
	t := []int{1,2,3,4,5}
	fmt.Println(t)

	// the slice package contain some usefull utility function
	t2 :=[]int{1,2,3,4,5}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// we also can make multi dimensional slice/array
	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("TwoD slice: ", twoD)

}
