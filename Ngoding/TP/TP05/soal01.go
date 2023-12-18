package main

import "fmt"

func inputTglPinjam_1303223079(tanggal, bulan, tahun *int) {
/* I.S. Menerima input nilai untuk tanggal, bulan, tahun
F.S. tanggal, bulan, tahun telah memiliki nilai*/
	var tgl, bln, thn int
	fmt.Scan(&tgl, &bln, &thn)
	*tanggal = tgl
	*bulan = bln
	*tahun = thn
	if !valid_1303223079(*tanggal, *bulan, *tahun) {
		fmt.Println("input tidak valid")
	}
}

func valid_1303223079(tanggal, bulan, tahun int) bool {
	var jumlahhri int
	getJumlahHari_1303223079(bulan, tahun, &jumlahhri)
	return tahun > 0 && bulan >= 1 && bulan <= 12 && tanggal > 0 && tanggal <= jumlahhri
}

func kabisat_1303223079(tahun int) bool {
	return tahun % 4 == 0
}

func getJumlahHari_1303223079(bulan, tahun int, jmlHari *int) {
/* I.S. Diketahui bulan dan tahun
F.S. jmlHari adalah jumlah hari dari bulan dan tahun tersebut*/
	if kabisat_1303223079(tahun) && bulan == 2 {
		*jmlHari = 29
	} else if bulan == 2 {
		*jmlHari = 28
	} else if bulan == 4 || bulan == 6 || bulan == 9 || bulan == 11{
		*jmlHari = 30
	} else{
		*jmlHari = 31
	} 
}

func hitungTanggalKembali_1303223079(tanggal1, bulan1, tahun1 int, tanggal2, bulan2, tahun2 *int) {
/* I.S. tanggal1, bulan1, tahun2 valid
F.S. tanggal2, bulan2, tahun2 adalah 3 hari berikutnya dari tanggal1, bulan1, tahun1*/
	var jmlhari int
	getJumlahHari_1303223079(bulan1, tahun1, &jmlhari)
	*bulan2 = bulan1
	*tahun2 = tahun1
	*tanggal2 = tanggal1 + 3
	if *tanggal2 > jmlhari {
		*tanggal2 = *tanggal2 - jmlhari
		*bulan2 = bulan1 + 1
		if *bulan2 > 12 {
			*bulan2 = *bulan2 - 12
			*tahun2 = tahun1 + 1
		}
	}
}

func main() {
	var tgl1, bln1, thn1 int
	var tgl2, bln2, thn2 int
	inputTglPinjam_1303223079(&tgl1, &bln1, &thn1)
	for valid_1303223079(tgl1, bln1, thn1) {
		hitungTanggalKembali_1303223079(tgl1, bln1, thn1, &tgl2, &bln2, &thn2)
		fmt.Println(tgl2, bln2, thn2)
		inputTglPinjam_1303223079(&tgl1, &bln1, &thn1)
	}

}