package main

import "fmt"

func main() {
	dict := map[string]string{
		"1": "Hello",
		"2": "World",
		"3": "!",
	}

	dict2 := map[int][]int{
		1: {1, 2, 3, 4},
		2: {10, 11, 12, 13},
	}

	fmt.Println(dict)
	fmt.Println(dict2)

	for _, k := range dict2 {
		fmt.Println(k)
		for _, i := range k {
			fmt.Println(i)
		}

	}

}
