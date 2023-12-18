package main

import "fmt"

type kartu [52]int

func main() {
	var N kartu
	var i, j int

	fmt.Scan(&N[i])
	i++

	for N[i-1] != 0 && i <= 52 {
		fmt.Scan(&N[i])
		i++
	}

	for j = i-2; j >= N[i+1]; j-- {
		fmt.Print(N[j], " ")
	}
}