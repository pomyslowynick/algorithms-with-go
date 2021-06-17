package main

import "fmt"

func main() {
	fmt.Println(GetNthFibonacci(12))
}

func GetNthFibonacci(n uint) uint {
	if n == 0 {
		return 0
	}

	var n1, n2 uint = 0, 1

	for i := uint(1); i < n; i++ {
		n1, n2 = n2, n1+n2
	}

	return n2
}
