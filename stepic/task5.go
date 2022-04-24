/* Последовательность состоит из натуральных чисел и завершается числом 0. Определите количество элементов этой последовательности, которые равны ее наибольшему элементу.

Формат входных данных

Вводится непустая последовательность натуральных чисел, оканчивающаяся числом 0 (само число 0 в последовательность не входит, а служит как признак ее окончания).

Формат выходных данных

Выведите ответ на задачу.

Sample Input:

1
3
3
1
0
Sample Output:

2*/
package main

import "fmt"

func main() {
	var a, count int
	count = 1
	element := -1

	for true {

		fmt.Scan(&a)
		if a == 0 {
			fmt.Println(count)
			break
		}
		switch {
		case a > element:
			element = a
			count = 1
		case a == element:
			count++
		default:
			continue
		}
	}

}
