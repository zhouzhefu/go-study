package main

import (
	"fmt" 
	"runtime"
	"time"
)

func tellOS() {
	fmt.Print("Go runs on ")
	switch os:=runtime.GOOS; os {
	case "darwin": fmt.Println("OS X")
	case "linux": fmt.Println("Linux")
	default: fmt.Printf("%S. ", os)
	}
	//fmt.Println(os) //"os" is not accessible outside switch
}

func replaceIfElse() {
	switch {
	case time.Now().Hour() < 12: 
		fmt.Println("Good Morning!")
	case time.Now().Minute() > 1: 
		fmt.Println("Tick Tock!")
	default: 
		fmt.Println("You win!")
	}
}


func main() {
	tellOS()
	replaceIfElse()
}