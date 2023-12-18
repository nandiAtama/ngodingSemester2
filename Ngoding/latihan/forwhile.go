package main

import "fmt"

func main() {
	var inp, res int
	fmt.Scan(&inp)
	for inp != 0 {
		if inp > 0 {
			res += inp
			fmt.Scan(&inp)
		} else if inp < 0 {
			res += inp
			fmt.Scan(&inp)
		}
	}
	fmt.Println(res)
}