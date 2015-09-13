package main

import "fmt"

func multiResult() (int, string) {
	return 30, "Ready?"
}

func namedReturn() (param1, param2 int) {
	param1 = 10
	param2 = 20
	return //"return" followed by nothing
} 

func main() {
	a, b := multiResult()
	fmt.Println("Int result is: ", a)
	fmt.Println("String result is: ", b)
	
	c, d := namedReturn()
	fmt.Println("First return is: ", c)
	fmt.Println("Second return is: ", d)
}