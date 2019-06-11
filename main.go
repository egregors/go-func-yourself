package main

import (
	"log"
)

// WTF?
//
// 1. no loops
// 2. no IFs
// 3. func is a single return
// 4. no side-effects
// 5. no assignments in funcs
// 6. only funcs with zero or one arguments
// 7. no arrays

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
	log.Printf("%d + %d = %d", 1, 3, add2(1, 3))
	log.Printf("%d + %d = %d\n", 1, 3, add1(1)(3))
	log.Printf("%d + %d = %d\n\n", 1, 3, add(1)(3))

	// -> no arrays
	// we need a lists
	log.Println("no arrays")
	log.Println("=========")
	log.Println("...so, that mean we need a pairs")

	// lets make simple pair
	p := pair("Answer")(42)
	log.Println(p)
	p1 := pair("Question")(p)
	log.Println(p1)
	log.Println("")

	// looks good, next we need funcions to get first or second element of pair
	// usually it call's fst | snd

	// get first or second element
	log.Println("fst: ", fst(p1))
	log.Println("snd: ", snd(p1))

	// let's make list of 3 elements
	p3 := pair(3)(pair("two")(pair(1)(nil)))
	log.Println(p3)
	log.Println("")

	// now we need a bridge from func world to imperative world
	// without func limitations

	log.Println("list -> arr")
	log.Println(list2arr(p3))

	p4 := arr2list([]interface{}{1, 2, 3})
	log.Println(p4)
	log.Println("")

	l5 := list2arr(arr2list([]interface{}{1, 2, 3, 4, 5}))
	log.Println(l5)
	log.Println("")

	// next
	// let's solve classic Fizz Buzz problem with all of this
	// n % 3 == 0 -> Fizz
	// n % 5 == 0 -> Buzz
	// n % 15 == 0 -> FizzBuzz
	//
	// [1, 2, 3, 4, 5, 6, 7, 8, 9, 10] ->
	// [1, 2, Fizz, 4, 5, Fizz, 7, 8, Fizz, 10]

	// first
	// we need kind of generator of range

	genList := gen(1)(100).(Pair)
	log.Println(list2arr(genList))

	// now we need map
	// T_T
	m := Map(func(el interface{}) interface{} { return el.(int) * el.(int) })(gen(1)(10).(Pair)).(Pair)
	log.Println(list2arr(m))
	log.Println("")

	log.Println("fizzbuzz as a single func")
	log.Println(fizzbuzz(1))
	log.Println(fizzbuzz(3))
	log.Println(fizzbuzz(5))
	log.Println(fizzbuzz(15))

	log.Println("aaaand the solve for problem in pure func way")

	log.Println(list2arr(IntStrMap(fizzbuzz)(gen(1)(100).(Pair)).(Pair)))
}
