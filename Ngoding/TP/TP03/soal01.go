package main

import (
	"fmt"
	"math"
)

func z(x, y int) float64 {
	var hasil float64
	hasil = math.Sqrt(float64(3*x) * float64(6*y) / float64(4*y))
	return hasil
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Printf("%2.3f %2.3f", z(a, b), z(b, a))
}