package main

import "fmt"

/*
* Unlike Java, Go has no Object, so there's no "ref" concept, and it is just 
* the reason why the infamous pointer is used. But compare to the pointer in 
* C lang, Go's pointer is more like Java's ref, since it does not allow 
* pointer arithmetic, and has GC to release the pointer address
*/
func assignVal() {
	i := 21
	p := &i
	fmt.Println(p) //just a memory address value
	//p = 3 //"p" is pointer so cannot be assigned a normal integer
	*p = *p/3 //value of "i" got changed
	fmt.Println(*p, i) //*p and i are actually the same thing
}

func main() {
	assignVal()
}