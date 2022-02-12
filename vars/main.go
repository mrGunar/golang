package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := "Andrei"
	const age int = 20

	message := ""
	message = "Hi from Go!"

	fmt.Println("Hi from Go!")
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(name, age)
	fmt.Println()
	fmt.Println(message)
	fmt.Println(reflect.TypeOf(name))
}
