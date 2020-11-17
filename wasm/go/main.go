package main

import (
	"fmt"
	"math"
	"syscall/js"
)

//Fibonacci algorithm
func Fibonacci(n uint) uint {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func isPrime(n int) bool {
	limit := int(math.Floor(math.Sqrt(float64(n))))
	for i := 2; i <= limit; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	c := make(chan bool)

	fmt.Println("Hello from Go")
	js.Global().Set("wsFibonacci", js.FuncOf(jsFibonacci))
	js.Global().Set("wsPrime", js.FuncOf(jsIsPrime))

	<-c
}

func jsFibonacci(this js.Value, args []js.Value) interface{} {
	input := args[0].Int()
	return Fibonacci(uint(input))
}

func jsIsPrime(this js.Value, args []js.Value) interface{} {
	input := args[0].Int()
	return isPrime(input)
}
