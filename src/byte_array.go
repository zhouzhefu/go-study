package main

import (
	"fmt"
	"crypto/md5"
	"io"
	"os"
)


func main() {
	/*
	bytes1 := []byte{53, 54, 52, 50, 51, 50, 54, 212, 29, 140, 217, 143, 0, 178, 4, 233, 128, 9, 152, 236, 248, 66, 126}
	fmt.Println(bytes1)
	fmt.Println(string(bytes1))
	
	bytes2 := []byte{25, 27, 244, 217, 128, 26, 104, 209, 142, 206, 225, 61, 224, 35, 194, 25, 34, 163, 48, 192, 204, 109, 175, 139, 249, 236, 223, 153, 231, 28, 80, 164}
	fmt.Println(bytes2)
	fmt.Sprintf("%s", bytes2)
	*/
	
	h := md5.New()
	c := h.Sum([]byte("test"))
	fmt.Sprintf("%x", []byte("test"))
	fmt.Println([]byte("test"))
	fmt.Println(string([]byte("test")))
	fmt.Println(fmt.Sprintf("%x", c))
	fmt.Println(c)
	fmt.Println(string(c))
	
	h1 := md5.New()
	io.WriteString(h1, "test")
	c2 := h1.Sum(nil)
	fmt.Sprintf("%x", c2)
	// fmt.Println(string(c2))
	
	bytes := []byte{53, 53, 53, 53}
	c1 := h.Sum(bytes)
	fmt.Sprintf("%x", c1)
 	// fmt.Println(c1)
	
	
	// writeResult(c)
}

func writeResult(resultBytes []byte) {
	file, err := os.OpenFile(
		"md5pass.txt", 
		os.O_WRONLY|os.O_CREATE, 
		os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	file.Write(resultBytes)
	file.Write([]byte("\n"))
	file.WriteString(string(resultBytes))
	file.Write([]byte("\n"))
	file.WriteString(fmt.Sprintf("%x", resultBytes))
	
}