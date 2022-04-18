package main

import "fmt"

func main() {
	fmt.Println("Hello")
	text := "Spam"

	result := text + fmt.Sprintf(" %d", 10) // Примерно как f-string Py

	fmt.Println(result)

	fmt.Println(fmt.Sprintf("%s %d %s", text, 50+4, result))
	fmt.Println(fmt.Sprintf("%T %T %T", text, 50, result))
	fmt.Println(fmt.Sprintf("%T %T %T", text, float32(50), result))

	var input float64
	fmt.Scanf("%f", &input)
	ouput := input + 1

	fmt.Println(ouput)
}
