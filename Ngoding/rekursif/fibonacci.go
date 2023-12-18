package main
import "fmt"

func main() {
	var cek int

	fmt.Scan(&cek)
	hasil:= s(cek)
	fmt.Println(hasil)
}

func s(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	} else {
		return s(n-2) + s(n-1)
	}
}