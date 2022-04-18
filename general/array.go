package main

import (
	"fmt"
	"sort"
)

func main() {
	var array [3]string
	list := [3]int{1, 2, 3} //static length

	fmt.Println(array)
	fmt.Println(list)
	fmt.Println("====================================")

	var slace []string //dinamic
	fmt.Println(slace)
	slace = append(slace, "foo")
	fmt.Println(slace[0])
	slace = append(slace, "spam")
	fmt.Println(slace)

	fmt.Println("================================")

	arr := make([]int, 10)
	for _, k := range arr {
		fmt.Println(k)
	}
	arr[4] = 20
	arr[2] = 42
	arr[len(arr)-1] = 100
	fmt.Println(arr)
	sort.Ints(arr)
	fmt.Println(arr)

}
