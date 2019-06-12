package main

import (
	"log"
	"strconv"
)

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

// very well
// if we got the list of elements, fst will return HEAD of the list and snd – TAIL of the list
var head = fst
var tail = snd

func gen(low int) func(int) interface{} {
	return func(high int) interface{} {
		if low > high {
			return nil
		}
		return pair(low)(gen(low + 1)(high))
	}
}

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

func Map(f func(interface{}) interface{}) func(Pair) interface{} {
	return func(xs Pair) interface{} {
		if tail(xs) == nil {
			return nil
		}
		return pair(f(head(xs)))(Map(f)(tail(xs).(Pair)))
	}
}

func IntStrMap(f func(int) string) func(Pair) interface{} {
	return func(xs Pair) interface{} {
		if tail(xs) == nil {
			return nil
		}
		return pair(f(head(xs).(int)))(IntStrMap(f)(tail(xs).(Pair)))
	}
}

func fizzbuzz(n int) string {
	// no assignment in functions? yep?
	s := map[bool]string{true: "Fizz", false: ""}[n%3 == 0] +
		map[bool]string{true: "Buzz", false: ""}[n%5 == 0]
	return map[bool]string{true: s, false: strconv.Itoa(n)}[s != ""]
}

func main() {
	res := list2arr(IntStrMap(fizzbuzz)(gen(1)(100).(Pair)).(Pair))
	log.Println(res)
}
