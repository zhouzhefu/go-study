package main

import (
	"fmt"
	"log"
	"net/http"
)

type HelloServerHandler struct{}
/*this is the "Handler" interface defined in "http" package*/
func (h HelloServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Web!")
	fmt.Println("Server echo!")
}

func runDefaultHandler() {
	var h HelloServerHandler
	//listen to "h" means use DefaultServeMux (as above "Hello Web!")
	err := http.ListenAndServe("localhost:4000", h)
	if (err != nil) {
		log.Fatal(err)
	}
}

/**
* Let's define 2 more Handler types
*/
type String string
func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, string(s))
}
type Struct struct{
	Greetings string
	Haha string
}
func (s *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func runAddedHandler() {
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello you", "hehe"})
	//listen to "nil" means use added handler, never hit DefaultServeMux
	err := http.ListenAndServe(":4000", nil)
	if (err != nil) {
		log.Fatal(err)
	}
}

func main() {
	//runDefaultHandler()
	runAddedHandler()
}

