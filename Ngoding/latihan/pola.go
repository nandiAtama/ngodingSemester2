package main
import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	hasil:= pola(n)
	fmt.Println(hasil)
}

func pola(bil int) string {
	var i, j int
	var hasil string
	hasil = ""
	for j=1; j<= bil; j++ {
		for i=1; i <= j; i++{
			hasil += "*"
		}
		hasil = hasil + "\n"
		
	}
	return hasil
}