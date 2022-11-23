package main

import (
	"fmt"
	"time"
)

func inputName() {
	var a int
	fmt.Println("Input number")
	fmt.Scan(&a)
	fmt.Println("Your input is: ", a)
}

func printNum() {
	for i := 0; i < 10; i++ {
		fmt.Print(i)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {

	go printNum()
	inputName()
	fmt.Printf("bye bye")
	time.Sleep(10 * time.Second)
}
