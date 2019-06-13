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

func main() {
	p1 := pair("Answer")(42)
	log.Println(p1)
	p2 := pair("Question")(p1)
	log.Println(p2)
}
