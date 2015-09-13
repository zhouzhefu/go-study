package main

import "fmt"


func range1() {
	pow := []int {1,3,5,7}
	for index, value := range pow {
		fmt.Println(index, value)
	}
}

func range2() {
	pow := []int {1,3,5,7}
	//we don't need the index here, so simply declare it as "_"
	for _, value := range pow {
		fmt.Println(value)
	}
	
	for index, _ := range pow {
		fmt.Println(index)
	}
	//equivalent as above, since we are ignoring the 2nd param, even a "_" 
	//is unnecessary
	for index := range pow {
		fmt.Println(index)
	}
	
	//or simply we just need a looping times, neither index or value needed
	//range() is just a function which returns 2 results: index, value
	for range(pow) {
		fmt.Println("OK")
	}
}

func main() {
	range1()
	range2()
}