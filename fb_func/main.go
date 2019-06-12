package main

import (
	"log"
	"strconv"
)

func fizzbuzz(n int) string {
	// no assignment in functions? yep?
	s := map[bool]string{true: "Fizz", false: ""}[n%3 == 0] +
		map[bool]string{true: "Buzz", false: ""}[n%5 == 0]
	return map[bool]string{true: s, false: strconv.Itoa(n)}[s != ""]
}

func main() {
	log.Println(fizzbuzz(1))
	log.Println(fizzbuzz(3))
	log.Println(fizzbuzz(5))
	log.Println(fizzbuzz(15))
}
