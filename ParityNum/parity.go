package main

import "fmt"

func isEvenMod(n int) bool {
	return n%2 == 0
}

func isEvenShift(n int) bool {
	return n>>1<<1 == n
}

func isEvenAnd(n int) bool {
	return n&1 == 0
}

func main() {
	var i int
	fmt.Scanf("%d", &i)
	isEvenMod(i)
	isEvenShift(i)
	isEvenAnd(i)
}
