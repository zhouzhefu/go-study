package main

import "fmt"

type Vertex struct {
	x int
	y int
}

func main() {
	v := Vertex{1,2}
	fmt.Println(v)
	
	v.x = 10
	fmt.Println(v)
	
	p := &v
	p.x = 20 //note when the pointer to struct access field, no * needed
	fmt.Println(*p)
}