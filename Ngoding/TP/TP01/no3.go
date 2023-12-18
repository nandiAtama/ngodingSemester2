package main

import (
	"fmt"
	"math"
)

func main() {
	var panjang, lebar, luas, keliling, diagonal float64

	fmt.Scan(&panjang, &lebar)
	luas = panjang * lebar
	keliling = panjang + lebar + panjang + lebar
	diagonal = math.Sqrt((panjang * panjang) + (lebar * lebar))
	fmt.Println("Luas: ", luas)
	fmt.Println("Keliling: ", keliling)
	fmt.Println("Panjang diagonal: ", diagonal)
}
