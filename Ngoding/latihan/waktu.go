package main

import "fmt"

func main() {
	 var jam, menit, detik int
	 fmt.Scan(&jam, &menit, &detik)
	 result := konversi(jam, menit, detik)
	 fmt.Println("Hasil konversi =", result, "detik")
}
func konversi(jam, menit, detik int) int {
	return (jam*3600)+(menit*60)+detik
}