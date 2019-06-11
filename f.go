package main

import "strconv"

// sum produce sum of sequence from 1 to n
func sum(n int) int {
	if n <= 0 {
		return 0
	}
	return n + sum(n-1)
}

// usual add
func add2(a, b int) int {
	return a + b
}

// func style add
func add1(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

// func style add, but prettier (no)
var add = func(a int) func(int) int { return func(b int) int { return a + b } }

// Pair is pair :-|
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

var arr2list = func(xs []interface{}) Pair {
	var res interface{}
	for i := len(xs) - 1; i >= 0; i-- {
		res = pair(xs[i])(res)
	}
	return res.(Pair)
}

func gen(low int) func(int) interface{} {
	return func(high int) interface{} {
		if low > high {
			return nil
		}
		return pair(low)(gen(low + 1)(high))
	}
}

// Map is map
func Map(f func(interface{}) interface{}) func(Pair) interface{} {
	return func(xs Pair) interface{} {
		if tail(xs) == nil {
			return nil
		}
		return pair(f(head(xs)))(Map(f)(tail(xs).(Pair)))
	}
}

// IntStrMap is map from ints to strs
func IntStrMap(f func(int) string) func(Pair) interface{} {
	return func(xs Pair) interface{} {
		if tail(xs) == nil {
			return nil
		}
		return pair(f(head(xs).(int)))(IntStrMap(f)(tail(xs).(Pair)))
	}
}

func fizzbuzz(n int) string {
	s := map[bool]string{true: "Fizz", false: ""}[n%3 == 0] + map[bool]string{true: "Buzz", false: ""}[n%5 == 0]
	return map[bool]string{true: s, false: strconv.Itoa(n)}[s != ""]
}
