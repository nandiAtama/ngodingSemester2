package main

import "fmt"

type ASEAN struct {
	negara                          string
	tahun, GDP                      int
	total2021, total2022, total2023 int
}
type arrNegara [100]ASEAN

func ASEAN_1303223079(T *arrNegara) {
	T[0].negara = "Brunei"
	T[5].negara = "Indonesia"
	T[1].negara = "Myanmar"
	T[6].negara = "Philippines"
	T[2].negara = "Cambodia"
	T[7].negara = "Singapore"
	T[3].negara = "Laos"
	T[8].negara = "Thailand"
	T[4].negara = "Malaysia"
	T[9].negara = "Vietnam"
}

func cariNegara_1303223079(T arrNegara, X string) int {

	for i := 0; i < 10; i++ {
		if X == T[i].negara {
			return i
		}
	}
	return -1
}

func input_1303223079(T *arrNegara) {
	var negara string
	var thn, gdp int

	fmt.Scanln(&negara, &thn, &gdp)
	i := cariNegara_1303223079(*T, negara)
	for i != -1 {
		T[i].tahun = thn
		T[i].GDP = gdp

		if T[i].tahun == 2021 {
			T[i].total2021 += T[i].GDP
		} else if T[i].tahun == 2022 {
			T[i].total2022 += T[i].GDP
		} else if T[i].tahun == 2023 {
			T[i].total2023 += T[i].GDP
		}
		fmt.Scanln(&negara, &thn, &gdp)
		i = cariNegara_1303223079(*T, negara)
	}
	sort_1303223079(T, thn)
}

func sort_1303223079(T *arrNegara, n int) {
	var pass, i int
	var temp ASEAN
	pass = 1

	for pass <= 10-1 {
		i = pass
		temp = T[pass]
		if n == 2021 {
			for i > 0 && temp.total2021 > T[i-1].total2021 {
				T[i] = T[i-1]
				i--
			}
		} else if n == 2022 {
			for i > 0 && temp.total2022 > T[i-1].total2022 {
				T[i] = T[i-1]
				i--
			}
		} else {
			for i > 0 && temp.total2023 > T[i-1].total2023 {
				T[i] = T[i-1]
				i--
			}
		}
		T[i] = temp
		pass++
	}
}

func main() {
	var tab arrNegara

	ASEAN_1303223079(&tab)
	input_1303223079(&tab)
	fmt.Println()

	for i := 0; i < 10; i++ {
		fmt.Println(tab[i].negara, tab[i].total2021, tab[i].total2022, tab[i].total2023)
	}
}
