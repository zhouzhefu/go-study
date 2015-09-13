package main

import "fmt"

var (
	str string //default value is empty string "", instead of null as in Java
	pack1, pack2 int //default value is 0 for type "int"
	pack3, pack4 = 2, 3 //with initializer, Go can defer by the initializer type
)

func main() {
	str += "a" //if the result is still "a", means the str init value is ""
	pack5, pack6 := 999, 9999 //shortcut to "var pack5, pack6 = 999, 9999"
	theFloat := float64(pack5) //convert int to float64
	fmt.Println(pack1, pack2, pack3, pack4, pack5, pack6, str, theFloat)
	printConst()
}

func printConst() {
	const BIG = 1 << 50
	const SMALL = BIG >> 49
	
	fmt.Println(BIG, SMALL)
}