package main

import "fmt"

func beliBuku_1303223079(x, y int) {
	var a int
	if y == y {
		a = x - y
	}
	if y != x {
		a = a - 1
		if a != 0 {
			fmt.Println("beli 1 buku, sisa", a)
			beliBuku_1303223079(x, y+1)
		} else if a == 0 {
			fmt.Println("beli 1 buku, rak buku penuh")
		}
	}
}
func main() {
	var x, y int
	fmt.Print("Kapasitas rak dan jumlah buku saat ini? ")
	fmt.Scanln(&x, &y)
	if x > y {
		fmt.Println("Sisa rak kosong", x-y, "buku")
		beliBuku_1303223079(x, y)
	}
}