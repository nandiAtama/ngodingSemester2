package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	display_1303223079(a, b)
}

func isFaktor_1303223079(num, x int) bool {
	return num%x == 0
}

func jFaktor_1303223079(num int, total *int) {
	*total = 0
	var i int = 1
	for i < num {
		if isFaktor_1303223079(num, i) {
			*total += i
		}
		i++
	}
}

func perfect_1303223127(num int) bool {
	var hasil int
	jFaktor_1303223079(num, &hasil)
	return hasil == num
}

func display_1303223079(x, y int) {
	var i int = x
	fmt.Print("Barisan Perfect Number: ")
	for i <= y {
		if perfect_1303223127(i) {
			fmt.Printf("%d", i)
		}
		i++
	}
}