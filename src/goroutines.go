package main

import (
	"fmt" 
	"time"
	
	"sync"

	"crypto/md5"
	"sort"
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

func inOutChan1() {
	var wg sync.WaitGroup
	totalCount := 6
	wg.Add(totalCount)
	
	// inChan := make(chan int, totalCount/2)
	inChan := make(chan int)
	
	inGo := func(in chan<- int) {
		for i:=0; i<totalCount; i++ {
			fmt.Println(i*100, "to be written.")
			in <- i*100
		}
		close(in)
	}
	go inGo(inChan)
	
	// this outGo() will read all the send-in values in the chan, therefore it 
	// is not possible to leave a writing gorouting hung forever. Otherwise, set 
	// the chan buffer size larger than or equals to the size of the writing set, 
	// so the writing goroutines will finish first without blocking. 
	outGo := func(out <-chan int) {
		for v := range out {
			fmt.Println(v, "was read!")
			time.Sleep(2 * time.Second)
			wg.Done() //sync counter minus 1
		}
	}
	
	go outGo(inChan)
	
	wg.Wait()
}

func inOutChan2() {
	var wg sync.WaitGroup
	
	merge := func(done <-chan struct{}, cs ...<-chan int) (chan int, sync.WaitGroup) {
		// wg.Add(len(cs)) //usually it is not a good idea
		
		out := make(chan int)
		output := func(c <-chan int) {
			// Increasing the counter. This must be done in main goroutine
			// because there is no guarantee that newly started goroutine will 
			// execute before 4 due to memory model guarantees.
			// wg.Add(1)
			defer wg.Done()
			
			for n := range c {
				fmt.Println("continue to merge: ", n)
				select {
				case out <- n:
					fmt.Println("Merged 'n':", n)
				case <-done:
					fmt.Println("Done signal received. Left:", len(done))
				// when "done" has items, do nothing but continue the loop, so 
				// the writing channel "c" will not be blocked due to "out" is 
				// blocking all input. 
				}
			}
			// wg.Done() //make it as "defer wg.Done()" at beginning is better
			fmt.Println("wg.Done()")
		}
		
		for _, c := range cs {
			// Increasing the counter. This must be done in main goroutine
			// because there is no guarantee that newly started goroutine will 
			// execute before 4 due to memory model guarantees.
			wg.Add(1) //Add(1) before start each new goroutine, much better
			// fmt.Println("add wg once, when it should be:", len(cs))
			go output(c)
		}
		
		return out, wg
	}
	
	
	sq := func(n int) chan int {
		startFrom := n
		out := make(chan int)
		
		wg.Add(1) //don't missed it for every child goroutine
		go func() {
			defer close(out)
			defer wg.Done()
			
			for i:=0; i<5; i++ {
				out <- n * n
				n += 1
			}
			fmt.Println("Square from:", startFrom, "writing is done.")
		}()
		
		return out
	}
	
	c1 := sq(1)
	c2 := sq(5)
	
	done := make(chan struct{}, 10)
	defer close(done)
	// for i:=0; i<8; i++ {
	// 	done <- struct{}{}
	// }
	out, wg := merge(done, c1, c2)
	fmt.Println("** Print 'out':", <-out)
	fmt.Println("** Print 'out':", <-out)
	
	
	for i:=0; i<10; i++ {
		done <- struct{}{}
	}
	
	// go func() {
	// 	fmt.Println("-- 'out' started to wait. ")
	// 	wg.Wait()
	// 	close(out)
	// 	fmt.Println("-- 'out' population is done & closed. ")
	// }()
	// time.Sleep(1 * time.Millisecond)
	
	// WHY previously try to Wait() at main goroutine always hit deadlock error? 
	// We forgot to use 'wg' to sync those goroutines started in sq(). So now 
	// please remember -- if you want to Wait() for exit the main goroutine, make 
	// sure ALL the child goroutines joining the WaitGroup, otherwise, you have 
	// to try your luck by putting Wait() in another child goroutine and set a 
	// sleep time at the end to pray those not joining WaitGroup child goroutines 
	// will exit before Wait() is being called. 
	fmt.Println("-- 'out' started to wait. ")
	wg.Wait()
	close(out)
	fmt.Println("-- 'out' population is done & closed. ")
}

/*
* Almost the same as inOutChan2(), the differences has been interpreted by comment
*/
func inOutChan3() {
	var wg sync.WaitGroup
	
	merge := func(done <-chan struct{}, cs ...<-chan int) (chan int, sync.WaitGroup) {		
		out := make(chan int)
		output := func(c <-chan int) {
			defer wg.Done()
			
			for n := range c {
				fmt.Println("continue to merge: ", n)
				select {
				case out <- n:
					fmt.Println("Merged 'n':", n)
				case <- done:
					// <- done will immediately return a zero-value when "done" 
					// chan was closed. 
					fmt.Println("Done signal received. Zero-value received:", <- done)
					return
				}
			}
			fmt.Println("wg.Done()")
		}
		
		for _, c := range cs {
			wg.Add(1)
			go output(c)
		}
		
		return out, wg
	}
	
	
	sq := func(n int) chan int {
		startFrom := n
		out := make(chan int)
		
		/*
		* sync seems no longer needed for generating goroutines: 
		* 1. Since traversing of this "out" chan is no longer guaranteed finished, 
		* 	 the "wg.Done()" may never be called, results in wg.Wait() may forever 
		* 	 blocked (the counter added here never get minus) because the main 
		* 	 goroutine want to process only a partial of the source chan, or 
		* 	 otherwise Go will hit fatal error of deadlock. 
		* 2. But there is a problem without sync, if the main goroutine happen to 
		* 	 trying traverse the whole source chan, this generating goroutines will 
		* 	 have no guarantee that the main goroutine will wait for them with 
		* 	 "wg.Wait()", deadlock fatal error could still be a problem. 
		*/
		// wg.Add(1)
		go func() {
			defer close(out)
			// defer wg.Done()
			
			for i:=0; i<5; i++ {
				out <- n * n
				n += 1
			}
			fmt.Println("Square from:", startFrom, "writing is done.")
		}()
		
		return out
	}
	
	c1 := sq(1)
	c2 := sq(5)
	
	// no more buffer assigned, since "done" exists just for a close()
	done := make(chan struct{})
	
	out, wg := merge(done, c1, c2)
	
	// fmt.Println("** Print 'out':", <-out)
	// fmt.Println("** Print 'out':", <-out)
	/* 
	* The loop size is a trap: 
	* 1. You don't try to traverse the whole "out" chan, or deadlock error will hit 
	* 	 due to forever blocking since "out" could never get closed before this line. 
	* 2. You must be carefully choose the loop size, making it no more than 
	* 	 maxLen(c1) + maxLen(c2), otherwise hit deadlock as described in point 1. 
	*/
	for i:=0; i<10; i++ {
		fmt.Println("** Print 'out':", <- out)
	}
	
	// after above "<-out", the program will be blocked until "done" get closed.  
	close(done)

	fmt.Println("-- 'out' started to wait. ")
	wg.Wait()
	close(out)
	fmt.Println("-- 'out' population is done & closed. ")
}

func basicGo() {
	go say("hello")
	//evaluation of params of goroutine body func will be done in the current 
	//goroutine, so no need to code like Java, that all params must be "final". 
	go say1(longTimedEvalParam())
	say("world")
	time.Sleep(2500 * time.Millisecond)
}

func serialDigestTree() {
	root := "/Users/winniewang/source/go-study/src"
	m, err := serialMD5All(root)
	if err != nil {
		fmt.Println(err)
		return
	}

	var paths []string
	for path := range m {
		paths = append(paths, path)
	}

	sort.Strings(paths)

	for _, path := range paths {
		fmt.Printf("%x %s\n", m[path], path)
	}
}

func serialMD5All(root string) (map[string][md5.Size]byte, error) {
	m := make(map[string][md5.Size]byte)
	err := nil

	if err != nil {
		return m, err
	}

	return m, nil
}

func main() {
	// basicGo()
	
	// inOutChan1()
	
	// inOutChan2()
	
	// inOutChan3()

	serialDigestTree()
}