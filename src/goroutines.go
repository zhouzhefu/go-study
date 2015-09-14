package main

import (
	"fmt" 
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func say1(s string) {
	time.Sleep(2000 * time.Millisecond)
	fmt.Println(s)
}

func longTimedEvalParam() string {
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("evaluation finished!")
	return "Hi"
}

func main() {
	go say("hello")
	//evaluation of params of goroutine body func will be done in the current 
	//goroutine, so no need to code like Java, that all params must be "final". 
	go say1(longTimedEvalParam())
	say("world")
	time.Sleep(2500 * time.Millisecond)
}