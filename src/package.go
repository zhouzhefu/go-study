package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(1000)
	fmt.Println("My Fav number is: ", rand.Intn(10))
}