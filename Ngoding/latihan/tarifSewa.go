package main

import "fmt"

func main() {
	var jam, menit int
	var member bool
	var total float64

	fmt.Scan(&jam, &menit, &member)

	hitungSewa_1303223079(jam, menit, member, &total)

	fmt.Println(total)
}

func durasi_1303223079(jam, menit int) int {
	if jam == 0 && menit != 0 {
		return 1
	} else if jam >= 1 {
		if menit >= 10 {
			return jam + 1
		} else {
			return jam
		}
	} else {
		return 0
	}
}

func potongan_1303223079(durasi, tarif int) float64 {
	if durasi > 3 {
		return float64(durasi*tarif) * 0.1
	} else {
		return 0
	}
}

func tarif_1303223079(member bool) int {
	if member {
		return 3500
	} else {
		return 5000
	}
}

func hitungSewa_1303223079(jam, menit int, member bool, biaya *float64) {
	var lama int = durasi_1303223079(jam, menit)
	var sewa int = tarif_1303223079(member)
	var diskon float64 = potongan_1303223079(lama, sewa)

	*biaya = float64(lama*sewa) - diskon
}