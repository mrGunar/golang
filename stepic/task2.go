/*
Дано неотрицательное целое число. Найдите и выведите первую цифру числа.

Формат входных данных
На вход дается натуральное число, не превосходящее 10000.

Формат выходных данных
Выведите одно целое число - первую цифру заданного числа.

Sample Input:

1234
Sample Output:
1
*/

package main

import "fmt"

func main() {
	var a int32
	fmt.Scan(&a)
	switch {
	case a == 10000:
		fmt.Println(1)
	case 1000 <= a && a < 10000:
		fmt.Println(a / 1000)
	case 100 <= a && a < 1000:
		fmt.Println(a / 100)
	case 10 <= a && a < 100:
		fmt.Println(a / 10)
	default:
		fmt.Println(a)
	}
}

func main2() {
	var s string
	fmt.Scan(&s)
	fmt.Print(string(s[0]))
}

func main3() {
	var a int
	fmt.Scan(&a)
	fmt.Println(getFirstNumber(a))
}

func getFirstNumber(a int) int {
	if a < 10 {
		return a
	} else {
		return getFirstNumber(a / 10)
	}
}
