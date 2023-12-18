package main

import "fmt"

type AsDos_T struct {
	kode, nama, mk    string
	nim, jumlah int
}

type TabAsDos_T[2500] AsDos_T

func bacaAsdos_1303223979(tabel *TabAsDos_T, n *int) {
	var i int

	i = 1

	for tabel[i-1].kode != "XXX" && i <= 2500 {
		fmt.Scanln(&tabel[i].kode, &tabel[i].nama, &tabel[i].nim, &tabel[i].mk, &tabel[i].jumlah)
		
		*n++
		i++
	}
}

func cetakAsdos_1303223079(tabel TabAsDos_T, n int, mk string) {
	var i int

	fmt.Println("Cetak Asdos MK : ")
	fmt.Scanln(&mk)

	for i = 0; i <= n; i++ {
		if tabel[i].mk == mk {
			fmt.Println("Asdos ke-", i)
			fmt.Println("Nama : ", tabel[i].nama)
			fmt.Println("Kode : ", tabel[i].kode)
			fmt.Println(" ")
		}
	}
}

func main() {
	var data TabAsDos_T
	var mk string
	var n int

	bacaAsdos_1303223979(&data, &n)
	cetakAsdos_1303223079(data, n, mk)

}