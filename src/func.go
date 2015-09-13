package main

import (
	"fmt"
	"math"
)

func funcAsVal() {
	f := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(f(3, 4))
}

/* 
* Fibonacci returns a function that returns successive Fibonacci numbers. 
* This can be taken as simulation of recurse func
*/ 
func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b //closure can use the bound variables a & b
		return a
	}
}

func main() {
	funcAsVal()
	
	f := fib()
	// Function calls are evaluated left-to-right.
	fmt.Println(f(), f(), f(), f(), f())
}																																																											