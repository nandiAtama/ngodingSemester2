package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func jarak_1303223079(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p1.x - p2.x, 2) + math.Pow(p1.y - p2.y, 2))
}

func main() {
	var A, B, C, D Point
	var JarakAB, JarakCD float64

	fmt.Scanln(&A.x, &A.y, &B.x, &B.y)
	JarakAB = jarak_1303223079(A, B)
	fmt.Scanln(&C.x, &C.y, &D.x, &D.y)
	JarakCD = jarak_1303223079(C, D)

	if JarakAB < JarakCD {
		fmt.Printf("Titik terdekat adalah titik A(%g,%g) dengan B(%g,%g) dengan jarak %g.", A.x, A.y, B.x, B.y, JarakAB)
	} else if JarakCD < JarakAB {
		fmt.Printf("Titik terdekat adalah titik C(%g,%g) dengan D(%g,%g) dengan jarak %g.", C.x, C.y, D.x, D.y, JarakCD)
	}

}