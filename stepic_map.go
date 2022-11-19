package main

import "fmt"

func main() {

	fmt.Println()
	// с помощью встроенной функции make:
	m1 := make(map[int]int)

	// с помощью использования литерала отображения:
	m2 := map[int]int{
		// Пары ключ:значение указываются при необходимости
		12: 2,
		1:  5,
	}
	fmt.Println(m1, "\n", m2)
	if value, inMap := m1[1]; inMap {
		fmt.Println(value) // 10
	}
}
