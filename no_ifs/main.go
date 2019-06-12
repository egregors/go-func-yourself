package main

import "log"

func sum(x int) int {
	s := make(map[bool]func(int) int)
	s[true] = func(y int) int { return y + sum(y-1) }
	s[false] = func(y int) int { return 0 }

	// o_o
	return s[x > 0](x)
}

func main() {
	// -> no IFs?
	// any ternary operators? no?
	// let's try make something like
	n := 42
	answer := map[bool]string{true: "negative", false: "positive"}[n < 0]
	log.Printf("%d is %s\n", n, answer)

	//	what about recursion?
	res := sum(5)
	log.Printf("sum(5) = %d", res)
}
