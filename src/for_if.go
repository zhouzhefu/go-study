package main

import (
	"fmt"
	"math"
)

func normalFor() {
	for i:=0; i<5; i++ {
		fmt.Print(i, ", ")
	}
	fmt.Println()
}

/*
* Go has no "while" keyword, for will do everthing about loop
*/
func whileFor() {
	i := 5
	for i > 0 {
		i -= 1
		fmt.Print(i, ", ")
	}
	fmt.Println()
}

func spellAAA() string {
	v := ""
	for i:=0; i<2; i++ {
		v += "A"
	}
	return v
}

func ifelse() string {
	u := "hello"
	if v:=spellAAA(); v == "AAA" {
		return v
	} else {
		return v+v //"v" can only be accessible within "if" block
	}
	
	return u
}

/* simulate the calculation of squareroot by using Newton's method
*/
func Sqrt(x float64) float64 {
	z := 1.0 //magical Newton, this "z" as seed can be any value!
	for i:=0; i<10; i++ {
		z = z - (z*z - x)/(2*z)
	}
	return z
}

func main() {
	normalFor()
	whileFor()
	fmt.Println(ifelse())
	
	fmt.Println(Sqrt(10), math.Sqrt(10))
}