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

func main() {
	p := pair("Answer")(42)
	log.Println(p)
	p1 := pair("Question")(p)
	log.Println(p1)
	log.Println("")

	// looks good, next we need functions to get first or second element of pair
	// usually it call's fst | snd
	// or head and tail

	// get first or second element
	log.Println("fst: ", head(p1))
	log.Println("snd: ", tail(p1))
}
