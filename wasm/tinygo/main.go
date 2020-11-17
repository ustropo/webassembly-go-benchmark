package main

import (
	"fmt"
	"math"
)

//export wsTinyFibonacci
func wsTinyFibonacci(n uint) uint {
	if n <= 1 {
		return n
	}
	return wsTinyFibonacci(n-1) + wsTinyFibonacci(n-2)
}

//export wsTinyIsPrime
func wsTinyIsPrime(n int) bool {
	limit := int(math.Floor(math.Sqrt(float64(n))))
	for i := 2; i <= limit; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Hello from TinyGo")
}
