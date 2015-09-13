package main

import "fmt"
import "strconv"
import "reflect"

func makingSlices() {
	a := make([]int, 5) //set length
	printSlice("a", a)
	
	b := make([]int, 0, 5) //In Java: new ArrayList<Integer>(5)
	printSlice("b", b)
	c := b[:cap(b)]
	printSlice("c", c)
	d := c[2:]
	printSlice("d", d)
	
	e := make([]int, 0, 0)
	if (e == nil) { //Nil array has len() and cap() 0, but reverse is not necessarily true, so this block could never be reached. 
		fmt.Println("e is Nil", len(e), cap(e))
	}
	var f []int
	if (f == nil) { 
		fmt.Println("f is Nil", len(f), cap(f))
	}
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

/*
* slice to array is basically ArrayList to Array in Java, providing much more 
* powerful functions and operations. And, more importantly, when the backend 
* array has not enough capacity to support current slice operation, a new larger 
* array will be allocated as the backend array of that slice. 
*/
func addElem() {
	a := [3]int {1, 3, 5} //with the given length "3", a is array
	s := []int {2, 4, 6, 8} //without the length given, s is slice
	fmt.Println(reflect.TypeOf(a), len(a), cap(a))
	fmt.Println(reflect.TypeOf(s), len(s), cap(s))
	
	//after append, it becomes a NEW slice
	s = append(s, 10)
	fmt.Println(s, len(s), cap(s))
}

func modifyArray(a1 [3]int) {
	a1[1] = 100
}

func modifySlice(s1 []int) {
	s1[1] = 100
}

func main() {
	//the Go array need not be initialized to have a zero value for each elem
	var a [10]string 
	
	for i:=0; i<10; i++ {
		//convert int to string, cannot as simple as Java way using "+"
		a[i] = strconv.Itoa(i) 
	}
	
	fmt.Println(a)
	
	b := []int{1, 3, 5, 7, 9} //note, "b" is a Slice rather Array!
	fmt.Println(b)
	
	s := b[2:4]
	fmt.Println(s)
	fmt.Println(s[1])
	s[1] = 13 //when change value of the slice, the backed array is changed 
	fmt.Println(b)
	
	makingSlices()
	
	addElem()
	
	//array in Go is a value? Let's try
	a1 := [3]int {3,5,7}
	fmt.Println("a1: ", a1)
	modifyArray(a1) //since array is "value", no side-effect possible
	fmt.Println("a1: ", a1) 
	//let's see what happen to slice
	s1 := a1[:]
	fmt.Println("s1: ", s1)
	modifySlice(s1) //ok now it works like Java ArrayList
	fmt.Println("s1: ", s1) //side-effect detected
	/*
	* so what is the conclusion? 
	* 1. array is "value"
	* 2. slice is "reference"
	* 3. you want to have side-effect, choose slice, otherwise array. 
	*/
}