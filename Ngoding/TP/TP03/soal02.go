package main
import "fmt"

func jumlahBus(penumpang, kapasitas int) int {
	var hasil int
	if penumpang % kapasitas == 0 {
		hasil = penumpang / kapasitas
	} else {
		hasil = penumpang / kapasitas + 1
	}
	return hasil
}

func main() {
	var hasil, n, m int
	fmt.Scan(&n, &m)
	hasil = jumlahBus(n, m) 
	fmt.Println(hasil, "bus")
}