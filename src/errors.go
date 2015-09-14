package main

import (
	"fmt"
	"strconv"
	"time"
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
//"type error interface {Error() string}" 
type MyError struct {
	When time.Time
	What string
}
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}
/**
* After method Error() link to MyError, we just finished a custom Exception 
* definition as we did in Java. 
*/

/**
* Note the Go has no "throws" keyword for exception, since Go func can return 
* multiple results, exception can be treated as a normal return result. 
* Please note the "error" is the built-in interface type
*/
func letMeHitError() error {
	fmt.Println("Let's ready to throw exception!")
	return &MyError{
		time.Now(), 
		"Just didn't work.", 
	}
}


func main() {
	fmt.Println(Vertex{3, 4})
	
	if err:=letMeHitError(); err != nil { //try..catch becomes if..else
		fmt.Println("Error catched: ", err)
	} else {
		fmt.Println("Everything going to be alright.")
	}
}