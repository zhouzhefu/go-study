package main

import (
	"fmt"
	"math"
)

/**
* To use interface we need a class first, and then let the class implement the 
* interface. But in Go there is no class, we use "duck model" instead. Therefore, 
* between the interface methods and the implementation methods, there is totally 
* no linking mechanism, as long as "duck" match, implementation is there. 
*/

//first of all we define an interface, 
//an interface is nothing but a collection of methods. 
type Abser interface {
	Abs() float64
}

//now we have 2 candidates (types) for this interface 
type Vertex struct {
	X, Y float64
}
type MyFloat float64 //yes, it is still a type

//Implement interface for Vertex, use pointer way. 
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}
//Implement interface for MyFloat, use value copy way. 
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	var ifce Abser
	
	ifce = &Vertex{3, 4}
	fmt.Println(ifce.Abs())
	
	ifce = MyFloat(-math.Pi)
	fmt.Println(ifce.Abs())
}