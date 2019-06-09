package main

import "log"

// WTF?
//
// 1. no loops
// 2. no IFs
// 3. func is a single return
// 4. no side-effects
// 5. no assignments in funcs
// 6. only funcs with zero or one arguments
// 7. no arrays

// sum produce sum of sequence from 1 to n
func sum(n int) int {
	if n <= 0 {
		return 0
	}
	return n + sum(n-1)
}

// usual add
func add_(a, b int) int {
	return a + b
}

// func style add
func add(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

func main() {
	// -> no loops?
	// recursion!
	log.Println("Sum")
	log.Println("===")
	log.Printf("1+2+3+4+5 = %d\n\n", sum(5))

	// -> no IFs?
	// any ternary operators? no?
	// let's try make something like
	//
	// n < 0 ? "negative" : "positive"
	log.Println("ternary")
	log.Println("=======")
	n := 42
	answer := map[bool]string{true: "negative", false: "positive"}[n < 0]
	log.Printf("%d is %s\n", n, answer)
	n = -69
	answer = map[bool]string{true: "negative", false: "positive"}[n < 0]
	log.Printf("%d is %s\n\n", n, answer)
	// but, obviously it will not be work with func.
	// both part of expression will be calculated in any cases
	//
	// 		func sum_(n int) int {
	// 			return map[bool]int{true: 0, false: n + sum_(n-1)}[n <= 0]
	// 		}

	// -> func is a single return?
	// without ternary operators it haven't many sense

	// -> no side-effects
	// only pure func exclude log's

	// -> no assignments in funcs
	// easy!

	// -> only funcs with zero or one arguments
	log.Println("only funcs with zero or one arguments")
	log.Println("=====================================")
	log.Printf("%d + %d = %d", 1, 3, add_(1, 3))
	log.Printf("%d + %d = %d\n\n", 1, 3, add(1)(3))

	// -> no arrays
	// we need a lists

}
