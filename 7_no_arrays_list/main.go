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

func main() {
	// let's make list of 3 elements
	p3 := pair(3)(pair("two")(pair(1)(nil)))
	log.Println(p3)
	log.Println("")

	log.Println("list -> arr")
	log.Println(list2arr(p3))

	p4 := arr2list([]interface{}{1, 2, 3})
	log.Println(p4)
	log.Println("")

	l5 := list2arr(arr2list([]interface{}{1, 2, 3, 4, 5}))
	log.Println(l5)
	log.Println("")
}
