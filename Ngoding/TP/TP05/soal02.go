package main

import "fmt"

func membeliKain(uangAwal float64, tUang, tPengeluaran *float64) {
	var temp float64
	temp = uangAwal / 4
	*tUang = *tUang - temp
	*tPengeluaran = *tPengeluaran + temp
}
func membeliBenangJahit(uangAwal float64, tUang, tPengeluaran *float64) {
	var temp float64
	temp = uangAwal / 20
	*tUang = *tUang - temp
	*tPengeluaran = *tPengeluaran + temp
	
}
func membuatPolaBaju(uangAwal float64, tUang, tPengeluaran *float64) {
	var temp float64
	temp = uangAwal / 200
	*tUang = *tUang - temp
	*tPengeluaran = *tPengeluaran + temp
}
func menjahitBaju(uangAwal float64, tUang, tPengeluaran *float64) {
	var temp float64
	temp = uangAwal / 200
	*tUang = *tUang - temp
	*tPengeluaran = *tPengeluaran + temp
}
func mengemasBaju(uangAwal float64, tUang, tPengeluaran *float64) {
	var temp float64
	temp = (3*uangAwal) / 200
	*tUang = *tUang - temp
	*tPengeluaran = *tPengeluaran + temp
}
func mendistribusikan(uangAwal float64, tUang, tPemasukan, tPengeluaran *float64) {
	var temp float64
	temp = (3*uangAwal) / 200
	*tPengeluaran = *tPengeluaran + temp
	*tPemasukan = *tPemasukan + (uangAwal/2)
	*tUang += -temp + *tPemasukan
}

func main() {
	var uangAwal, tUang, tPengeluaran, tPemasukan float64
	fmt.Scan(&uangAwal)
	tUang = uangAwal
	membeliKain(uangAwal, &tUang, &tPengeluaran)
	membeliBenangJahit(uangAwal, &tUang, &tPengeluaran)
	membuatPolaBaju(uangAwal, &tUang, &tPengeluaran)
	membuatPolaBaju(uangAwal, &tUang, &tPengeluaran)
	menjahitBaju(uangAwal, &tUang, &tPengeluaran)
	menjahitBaju(uangAwal, &tUang, &tPengeluaran)
	menjahitBaju(uangAwal, &tUang, &tPengeluaran)
	menjahitBaju(uangAwal, &tUang, &tPengeluaran)
	mengemasBaju(uangAwal, &tUang, &tPengeluaran)
	mengemasBaju(uangAwal, &tUang, &tPengeluaran)
	mendistribusikan(uangAwal, &tUang, &tPemasukan, &tPengeluaran)
	fmt.Printf("%.0f %.0f %.0f", tUang, tPemasukan, tUang)













//   var x float64
//   fmt.Scan(&x)
//   pengeluaran = (1 * x + membeliKain(uangAwal, tUang, tPengeluaran) + 1 * x + membeliBenangJahit(uangAwal, tUang, tPengeluaran) + 2 * x + membuatPolaBaju(uangAwal, tUang, tPengeluaran) + 4 * x + menjahitBaju(uangAwal, tUang, tPengeluaran) + 2 * x + mengemasBaju(uangAwal, tUang, tPengeluaran))
//   pemasukan = (1 * x + mendistribusikan(uangAwal, tUang, tPengeluaran))
//   uangJojo = pengeluaran - pemasukan
//   fmt.Println(pengeluaran, pemasukan, uangJojo)
}