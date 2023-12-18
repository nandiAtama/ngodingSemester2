package main 

import "fmt"

func main() {
	var tahun, totalBelanja, A, CD, diskon, dibayar int

	fmt.Scan(&tahun, &totalBelanja)
	A = tahun / 1000
	CD = tahun % 100
	diskon = A * CD
	dibayar = totalBelanja - (totalBelanja*diskon)/100
	fmt.Println("besar diskon: ", diskon, "%")
	fmt.Println("Jumlah yang dibayar: ", dibayar)
}
