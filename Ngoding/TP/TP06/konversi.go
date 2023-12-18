package main

import "fmt"

func abc_1303223079(x int) string {
	var a int
	var s string
	s = ""
	if x != 0 {
		a = x % 2
		x = x / 2
		if a == 0 {
			s = abc_1303223079(x) + "0"
		} else {
			s = abc_1303223079(x) + "1"
		}
	} else {
		return s
	}
	return s
}

func main() {
	var i int
	fmt.Scan(&i)
	fmt.Print(abc_1303223079(i))