package main

import "fmt"

func main() {
	var golongan string
	var jamKerja, lemburan, gajiA, gajiB, totalGaji int
	gajiA = 0
	gajiB = 0
	totalGaji = 0

	for  golongan != "Z" {
		fmt.Scan(&golongan, &jamKerja, & lemburan)
		if golongan == "A" {
			gajiA = (jamKerja * 75000) + (lemburan * 90000)
			fmt.Println(gajiA)
		} else if golongan == "B" {
			gajiB = (jamKerja * 125000) + (lemburan * 70000)
			fmt.Println(gajiB)
		} else if golongan == "Z" {
			totalGaji = gajiA + gajiB
		} else  {
			fmt.Println("Golongan tidak dikenali")
		}
	}
	fmt.Println("Biaya yang dikeluarkan perusahaan untuk gaji karyawan ", totalGaji)
}