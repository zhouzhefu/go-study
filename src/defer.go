package main

import "fmt"

func main() {
	for i:=0; i<3; i++ { //defer calls were put into a LIFO stack
		defer fmt.Print(i, ", ")
	}
	
	fmt.Println("Ok now let's count it after return. ")
}