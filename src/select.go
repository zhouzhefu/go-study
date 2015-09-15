package main

import (
	"fmt"
	"time"
)

func main() {
	//tick & boom are of the type "<-chan". As "chan" is a bidirectional channel, 
	//"<-chan" and "chan<-" are one-directional channel. 
	tick := time.Tick(100 * time.Millisecond)
	boom := time.Tick(500 * time.Millisecond)
	
	for tickCnt := 0;; {
		select {
		case <- tick: 
			fmt.Println("tick.")
			tickCnt += 1
		case <- boom:
			fmt.Println("Boom!")
		default: 
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
		
		if tickCnt > 10 {
			break;
		}
	}
}