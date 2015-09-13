package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

// when several params share the type, only the last one needs to mark
func add2(x, y, z int) int {
	return x + y + z
}

/* 
* further explanation of add2(), the "shared type" starts from y only, 
* since the previous param "echo" has been marked the type
*/
func add3(echo string, y, z int) int {
	fmt.Println(echo)
	return y + z
}

func main() {
	fmt.Println(add(4, 7))
	fmt.Println(add2(2, 3, 5))
	fmt.Println(add3("Echo me", 8, 7))
}