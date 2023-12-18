package main

import "fmt"

func len(num int) int {
	var total int = 0
	for num > 0 {
		total++
		num = num / 10
	}
	return total
}

func pangkat_1303223127(n int) int {
	if n == 0 {
		return 1
	} else {
		return 10 * pangkat_1303223127(n-1)
	}
}

func split_1303223127(num int, num1, num2 *int) {
	var panjang int = len(num)
	var tengah int = panjang / 2

	var pembagi int = pangkat_1303223127(tengah)

	*num2 = num % pembagi
	*num1 = num / pembagi
}

func main() {
	var num, num1, num2 int
	fmt.Scanln(&num)

	split_1303223127(num, &num1, &num2)

	fmt.Println(num1, num2)
	fmt.Println(num1 + num2)
}