package main

import "log"

type Pair struct {
	a, b interface{}
}

var pair = func(a interface{}) func(interface{}) Pair {
	return func(b interface{}) Pair {
		return Pair{a, b}
	}
}

var fst = func(p Pair) interface{} { return p.a }
var snd = func(p Pair) interface{} { return p.b }

var head = fst
var tail = snd

var list2arr = func(xs Pair) []interface{} {
	var res []interface{}
	for {
		if tail(xs) == nil {
			return append(res, head(xs))
		}
		res = append(res, head(xs))
		xs = tail(xs).(Pair)
	}
}

func gen(low int) func(int) interface{} {
	return func(high int) interface{} {
		if low > high {
			return nil
		}
		return pair(low)(gen(low + 1)(high))
	}
}

func main() {
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
}
