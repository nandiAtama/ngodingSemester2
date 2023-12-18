package main

import "fmt"

func main() {
	var i, uang, jumlah, rataRata, uangTertinggi, uangTerendah int
	jumlah = 0
	i = 0

	fmt.Scan(&uang)
	uangTertinggi = uang
	uangTerendah = uang
	for uang > 0 {
		jumlah = jumlah + uang
		if uang > uangTertinggi  {
			uangTertinggi = uang
		} else if uang < uangTerendah {
			uangTerendah = uang
		} 
		i++
		rataRata = jumlah / i
		fmt.Scan(&uang)
	}
	fmt.Println("Jumlah =", jumlah)
	fmt.Println("Rata - rata =", rataRata)
	fmt.Println("Uang tertinggi =", uangTertinggi)
	fmt.Println("Uang terendah =", uangTerendah)
}
