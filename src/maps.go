package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func map1() {
	//Map must be "make" before use, similar to "var", if declare with 
	//initializer, "make" can be omitted
	m := make(map[string]Vertex) 
	
	m["bell"] = Vertex {
		30.398, 50.233,
	}
	m["ball"] = Vertex {25.334, 77.345}
	
	fmt.Println(m)
}

func map2() {
	//initialize Map with key-value pairs, it is a bit similar to 
	//array & struct literals
	m := map[string]Vertex {
		"bellLabs": Vertex {53.223, 56.333}, 
		"google": Vertex {189.334, 33.353},
	}
	
	fmt.Println(m)
	
	//you can even omit the type name in literals
	m1 := map[string]Vertex {
		"bellLabs": {50.223, 33.333}, 
		"google": {18.334, 133.353},
	}
	
	fmt.Println(m1)
}

func mutatingMap() {
	m := map[string]int {
		"number1": 1,
		"number2": 2,
		"number10": 10,
	}
	fmt.Println(m)
	
	//now remove one element
	delete(m, "number10")
	fmt.Println(m)
	
	//same as in Java: Map.containsKey() + Map.get(), actually the map[key] is 
	//just a func returns 2 results, usually when we use it in this way: 
	//elem = map[key], the 2nd result is not needed so we similar don't declare. 
	//Apparently, as a non-OO lang, multi-result func is the best way to simulate 
	//"overlaoding" in Java, just be careful to arrange the sequence of results.
	elem, exists := m["number2"]
	fmt.Println("Key [number2] exists?", exists, ". And value is:", elem)
	
	//zero value when key not found
	elem1, exists1 := m["dummy"]
	fmt.Println("Zero value of dummy key?", elem1, exists1)
}

func traverseMap() {
	m := map[string]int {
		"key1": 10, 
		"key2": 20, 
		"key3": 30, 
	}
	
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func main() {
	map1()
	map2()
	mutatingMap()
	
	traverseMap()
}