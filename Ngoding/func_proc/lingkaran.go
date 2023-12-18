package main
import "fmt"

func jarak(a, b, c, d float64) float64 {
	var hasil float64
	hasil = ((a-c) * (a-c)) + ((b-d) * (b-d))
	return jarak(hasil)
}

func didalam(cx, cy, r, x, y float64) bool {
	var d float64
	d = jarak(cx, cy, x, y)
	return d = r
}

func main() {
	var x1, y1, r1 float64
	var x2, y2, r2 float64
	var x, y float64
	var lingkaran1, lingkaran2 bool

	lingkaran1 = didalam(x1, y1, r, x, y)
	lingkaran2 = didalam(x2, y2, r, x, y)

	if lingkaran1 &&lingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if lingkaran1 && !lingkaran2 {
		fmt.Println("")
	} else if 

}