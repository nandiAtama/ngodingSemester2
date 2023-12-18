package main

import "fmt"

func cumlaude(ipk float64, masaStudi int, publikasi bool) bool {
	var hasil bool
	hasil = (3.51 <= ipk && ipk <= 4.00) && masaStudi <= 8 && publikasi
	return hasil
}

func sangatMemuaskan(ipk float64, masaStudi int, publikasi bool) bool {
	var hasil bool
	hasil = (3.51 <= ipk && ipk <= 4.00) || (2.76 <= ipk && ipk <= 3.5) && masaStudi <= 14 && publikasi
	return hasil
}

func memuaskan(ipk float64, masaStudi int, publikasi bool) bool {
	var hasil bool
	hasil = (2.00 <= ipk && ipk <= 2.75) && masaStudi <= 14 || publikasi
	return hasil
}

func main() {
	var ipk float64
	var semester int
	var publikasi bool
	fmt.Scan(&ipk, &semester, &publikasi)
	if cumlaude(ipk, semester, publikasi) {
		fmt.Print("Cumlaude")
	} else if sangatMemuaskan(ipk, semester, publikasi) {
		fmt.Print("Sangat Memuaskan")
	} else if memuaskan(ipk, semester, publikasi) {
		fmt.Print("Memuaskan")
	}

}