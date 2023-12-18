package main
import "fmt"

func zoomIn(skala int, panjang, lebar *int) {
	*panjang = *panjang * skala
	*lebar = *lebar * skala
}

func zoomOut(skala int, panjang, lebar *int) {
	*panjang = *panjang / skala
	*lebar = *lebar / skala
}

func main() {
	var panjang, lebar, skala int
	var s string

	fmt.Scan(&panjang, &lebar)
	fmt.Scan(&s, &skala)

	if s == "+" {
		zoomIn(skala, &panjang, &lebar)
	} else if s == "-" {
		zoomOut(skala, &panjang, &lebar)
	}
	fmt.Println(panjang, lebar)
}