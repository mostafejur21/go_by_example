package main

import (
	"fmt"
	"maps"
)

func main() {
	// NOTE: to create map, we need to use builtin method make:
	// make(map[key-type]value-type)
	m := make(map[string]int)

	// setting map is name[key] = val
	m["k1"] = 7
	m["k2"] = 19
	fmt.Println("map:", m)

	// we can assign the value of a variable
	v1 := m["k1"]
	fmt.Println(v1) // 7

	// if any key doesn't exist, then the zero value will be shown
	v2 := m["k3"]   // k3 key does not exist in the map
	fmt.Println(v2) // 0

	// the builtin len returns the number of key/value pairs in the map
	fmt.Println("len:", len(m))

	// builtin delete method delete the key/value pairs from a map
	delete(m, "k2")
	fmt.Println("map:", m)

	// to remove all key/values pair in a map, we need to use clear() func
	clear(m)

	// this is to check weather the key present at the map or not. this will return a bool
	_, prs := m["k2"]

	fmt.Println("prs:", prs)

	// we can also declare and initilize a new map at the same line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map n:", n)

	// map package contian very usefull utility functions for a maps
	m2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, m2) {
		fmt.Println("n == m2")
	}
}
