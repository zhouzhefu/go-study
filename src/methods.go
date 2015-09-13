package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

/**
* Go has no classes, then how to make a "method" whose "receiver" is a type?
*/
func methodReceiver() {
	v := &Vertex{3,4}
	fmt.Println(v.Pythagorean())
}

/* 
* method name is "Pythagorean", (v *Vertex) is to tell the receiver of this 
* method is a Vertex instance. Unfortunately this "type" cannot be defined 
* outside current package
* Upon experiment we found such method cannot be defined in another func
*/
func (v *Vertex) Pythagorean() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func receiverAsValue() {
	v := Vertex{3, 4}
	v1 := v.CopyScale()
	fmt.Println("Address of the v == v1:", &v == &v1)
	
	vp := &v
	v2 := vp.RefScale()
	fmt.Println("Address of the vp == v2:", vp == v2)
}

//param passed as reference
func (v *Vertex) RefScale() *Vertex {
	v.X *= 2
	v.Y *= 2
	return v
}
/*
* Param passed as value, every time the method called the param will be copied 
* once. So be careful to choose, if an instance of param is heavy to copy, 
* you may want to use pointer receiver when design the mehod. 
* 
* But a good thing for such method is, it is immutable & side-effect free, try 
* to use it as much as possible to satisfy a "supple design". 
*/
func (v Vertex) CopyScale() Vertex {
	v.X *= 2
	v.Y *= 2
	return v
}

func main() {
	methodReceiver()
	
	receiverAsValue()
}