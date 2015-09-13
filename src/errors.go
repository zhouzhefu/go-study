package main

import (
	"fmt"
	"strconv"
)

type Vertex struct {
	X, Y int
}
//If want to override toString() for a type like we did in Java, just add
//String() method to this type, making it implements the built-in interface 
//"Stringer". 
func (v Vertex) String() string {
	return "Vertex: " + strconv.Itoa(v.X) + "," + strconv.Itoa(v.Y)
}

//Similar to Stringer, error is also a built-in interface: 
//type error interface {Error() string}



func main() {
	fmt.Println(Vertex{3, 4})
}