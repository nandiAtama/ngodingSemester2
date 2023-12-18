package main
import "fmt"

func denom(uang int, k10, k2, k1 *int) {
	*k10 = uang / 100000
	*k2 = uang % 10000 / 2000
	*k1 = uang % 10000 % 2000 / 1000
}

func main() {
	var uang, puluh, dua, satu int
	fmt.Scan(&uang)
	denom(uang, &puluh, &dua, &satu)
	fmt.Println(puluh, "Lembar 10000")
	fmt.Println(dua, "Lembar 2000")
	fmt.Println(satu, "Lembar 1000")
}