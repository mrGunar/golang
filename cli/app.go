package main

import (
	"flag"
	"fmt"
)

// The first way
var name = flag.String("name", "Foo", "A name to say hello")

// The second way
var spanish bool

func init() {
	flag.BoolVar(&spanish, "spanish", false, "Spanish language")
	flag.BoolVar(&spanish, "s", false, "Spanish language")
}

func main() {
	flag.Parse()
	if spanish == true {
		fmt.Printf("Hola, %s!\n", *name)
	} else {
		fmt.Printf("Hello %s!", *name)

	}
}
