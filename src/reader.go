package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!") //use string as example
	b := make([]byte, 8)
	
	for { //while(true)
		n, err := r.Read(b) //"n" is the number of bytes read
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if (err == io.EOF) { //catch (IoEOFException err) {...}
			break
		}
	}
}