package main
import "fmt"

// func pola(n int, s *string) {
// 	if n == 1 {
// 		*s = "*"
// 	} else {
// 		pola(n-1, s)
// 		*s += "*"
// 	}
// 	fmt.Println(*s)
// }

func main() {
	var n int
	var s string 
	fmt.Scan(&n)
	pola(n, &s)
}

func pola(n int, s *string) {
	if n == 1 {
		*s = "*"
	} else {
		pola(n-1, s)
		*s += "*"
	}
	fmt.Println(*s)
}