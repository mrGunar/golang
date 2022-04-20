package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	f, v := getFiles(".")
	fmt.Println(f)
	fmt.Println(v)

	spam := func(x, y int) int {
		res := x + y
		return res
	}
	fmt.Println(spam(1, 2))
}

func getFiles(path string) ([]string, error) {
	var files []string
	dirFiles, err := ioutil.ReadDir(fmt.Sprintf("%s", path))
	if err != nil {
		return nil, err
	}
	for _, k := range dirFiles {
		files = append(files, k.Name())
	}
	return files, nil
}

// Можно задать сразу имя возвращаемой переменной
func foo() (vars []int) {
	vars = []int{1, 2, 3, 4, 5}
	return
}
