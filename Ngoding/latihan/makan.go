package main

import "fmt"

func main() {
	var menu, n, m, i, belanja int
	var sisa bool

	fmt.Scan(&m)
	for i = 0; i < m; i++ {
		fmt.Scan(&menu, &n, &sisa)
		hitungTarif_1303223127(menu, sisa, n, &belanja)
		fmt.Println(belanja)
	}
}

func tarif_1303223127(menu int) int {
	var tarif int
	if menu > 50 {
		tarif = 100000
		return tarif
	} else if menu > 3 {
		tarif = 10000 + ((menu - 3) * 2500)
		return tarif
	} else {
		tarif = 10000
		return tarif
	}
}

func hitungTarif_1303223127(menu int, bersisa bool, n int, total *int) {
	if bersisa {
		*total = n * tarif_1303223127(menu)
	} else {
		*total = tarif_1303223127(menu)
	}
}