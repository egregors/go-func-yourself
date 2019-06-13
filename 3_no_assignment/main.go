package main

import "log"

func sum(x int) int {
	return map[bool]func(int) int{
		true:  func(y int) int { return y + sum(y-1) },
		false: func(y int) int { return 0 },
	}[x > 0](x)
}

func main() {
	res := sum(5)
	log.Printf("sum(5) = %d", res)
}
