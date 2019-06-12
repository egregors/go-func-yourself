package main

import "log"

func sum(n int) int {
	if n <= 0 {
		return 0
	}
	return n + sum(n-1)
}

func main() {
	res := sum(5) // 1 + 2 + 3 + 4 + 5
	log.Printf("sum(5) = %d", res)
}
