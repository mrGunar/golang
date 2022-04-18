package main

import "fmt"

func main() {
	foo(1, 2, 3)
}

func foo(a, b, c int) {
	res := (a + b + c) / 3
	if res >= 90 && res <= 100 {
		fmt.Println("A")
	}
	if res >= 80 && res < 90 {
		fmt.Println("B")
	}
	if res >= 70 && res < 80 {
		fmt.Println("C")
	}
	if res >= 60 && res < 70 {
		fmt.Println("D")
	}
	if res >= 0 && res < 60 {
		fmt.Println("F")
	}
}

func spam(a, b, c int) string {
	switch mean := (a + b + c) / 3; {
	case mean < 60:
		return "F"
	case mean < 70:
		return "D"
	case mean < 80:
		return "C"
	case mean < 90:
		return "B"
	default:
		return "A"
	}
}
