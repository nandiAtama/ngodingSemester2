package main
import "fmt"

func volume(r, t float64) float64 {
	return pi * (r*r) * t
}

func massa(r, t, m float64) float64 {
	return volume(r, t) * m
}

func display(m1, m2 float64) {
	if m1 == m2 {
		fmt.Println("BALANCE")
	} else if m1 > m2 {
		fmt.Println(m1-m2)
	} else {
		fmt.Println(m2 - m1)
	}
}

const pi float64 = 3.14
func main() {

	var r, tKiri, mKiri, tKanan, mKanan float64
	var massaKiri, massaKanan float64

	fmt.Scan(&r)
	fmt.Scan(&tKiri, mKiri)

	massaKiri = massa(r, tKiri, mKiri)
	massaKanan = massa(r, tKanan, mKanan)

	display(massaKiri, massaKanan)

}