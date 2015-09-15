package main

import (
	"fmt"
	"time"
)

func sum(s []int, ch chan int) {
	//this will change the goroutine sequence to write channel
	if s[0] == 1 {
		time.Sleep(50 * time.Millisecond)
	} 
	
	sum := 0
	for _,v := range s {
		sum += v
	}
	ch <- sum
	
	fmt.Println("write channel done")
}

/**
* sometimes, channel works as BlockingQueue ...
*/
func testSumChannel() {
	src := []int {1,2,3,4,5,6,7,8,9,10}
	ch := make(chan int) //new BlockingQueue<Integer>()
	//ch := make(chan int, 100) //100 is the buffer size, same as BlockingQueue
	go sum(src[:len(src)/2], ch)
	go sum(src[(len(src)/2+1):], ch)
	
	//as a blockingQueue, "chan" is FIFO, and block writing until existing item 
	//received.  
	result1 := <- ch
	fmt.Println("result1 read!")
	result2 := <- ch
	fmt.Println("result2 read!")
	//result1, result2 := <- ch, <- ch //another style, but from this cannot tell 
										// the blocking mechanism
	
	fmt.Println(result1, result2)
}

/**
* and sometimes, channel works sd io.stream, or a netty style nio2.channel. 
* 
* You may found here the "chan" is passed to fibonacci() function by value, why?
* As interpreted in this blog, and in the referred Go FAQ: 
* http://my.oschina.net/chai2010/blog/161384
* map/slice/channel are "reference type" and the backend datastore is a pointer 
* field. E.g.: 
* 
* type ARefType *[10]int 
* 
* When using the ARefType, for example pass it to a func as param, what you see 
* is a value copied for evaluating func params, but the copied value is a pointer, 
* no matter how many times you copy the ref type, its backed array is never 
* destroyed or re-created. 
*/
func testCloseChannel() {
	fibonacci := func(cnt int, ch chan int) {
		x, y := 0, 1
		for i:=0; i<cnt; i++ {
			ch <- x
			x, y = y, x+y
		}
		close(ch)
	}
	
	ch := make(chan int, 10)
	
	go fibonacci(cap(ch), ch)
	//this is the special offer syntax sugar to read a channel until closed
	for fibVal := range ch {
		fmt.Println(fibVal)
	}
	//equivalently, the mechanism should be like this: 
	/*
	for {
		fibVal, ok := <- ch //if ok is "false", the channel was closed at here
		if (ok) {
			fmt.Println(fibVal)
		} else {
			fmt.Println("Read Channel Done!")
			break;
		}
	}
	*/
}

func main() {
	testSumChannel()
	
	testCloseChannel()
}