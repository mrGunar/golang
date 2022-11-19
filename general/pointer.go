package main

import "fmt"

func main() {
	var a int
	fmt.Print(&a) // -> 0xc000014078
}
