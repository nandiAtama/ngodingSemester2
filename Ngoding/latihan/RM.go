package main

import "fmt"

func tarif(makan int)int{
	var harga int
	harga = 10000
	if makan == 50{
		harga = 100000
	}else if makanan > 3{
		makan -= 3
		harga += makan * 2500
	}
	if makan-=50 > 0 {
		harga+=makan*2500
	}
	return harga
}

func hitungTarif(menu, orang int,isHabis bool, totalHarga *int){
	*totalHarga = tarif(menu)
	if !isHabis{
		// *totalHarga += (*totalHarga / orang) * orang
		*totalHarga += orang * 2000
	}
}

func main() {
	var rombongan, menu, orang, totalHarga int
	var isHabis bool
	fmt.Scan(&rombongan)
	for i := 0; i < rombongan; i++ {
		totalHarga = 0
		fmt.Scan(&menu,&orang,&isHabis)
		hitungTarif(menu,orang,isHabis,&totalHarga)
		fmt.Println(totalHarga)
	}
}