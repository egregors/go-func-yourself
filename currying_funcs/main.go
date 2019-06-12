package main

import "log"

// usual add
func add1(a, b int) int {
	return a + b
}

// func style add
func add2(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

// func style add, but prettier (no)
var add = func(a int) func(int) int { return func(b int) int { return a + b } }

func main() {
	res1 := add1(20, 19)
	res2 := add2(20)(19)
	res := add(20)(19)

	log.Printf(""+
		"add1: %d\n"+
		"add2: %d\n"+
		"add: %d\n",
		res1, res2, res)
}
