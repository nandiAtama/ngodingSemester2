package main

import "fmt"

func pola_1303223079(n int, s string) {
	if n != 0 {
		fmt.Println(s)
		s = "*" + s
		pola_1303223079(n-1, s)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	pola_1303223079(n, "*")
}