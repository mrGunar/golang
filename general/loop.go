package main

import "fmt"

func main() {
	i := 0
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for j := 10; j <= 20; j++ {
		fmt.Println(j)
	}

	text := []string{"a", "b", "c", "d"}

	for i, k := range text {
		fmt.Printf("%d%s ", i, k)
	}
}
