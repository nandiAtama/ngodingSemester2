package main
import "fmt"

// func ganjil(n int) {
// 	if n == 1{
// 		fmt.Print(1, " ")
// 	} else {
// 		ganjil(n-1)
// 		if !(n%2==0) {
// 			fmt.Print(n, " ")
// 		}
// 	}
// }

func main() {
	var n int
	fmt.Scan(&n)
	ganjil(n)
}

func ganjil(n int) {
	if n == 1 {
		fmt.Print(1, " ")
	} else {
		ganjil(n-1)
		if n%2==1 {
			fmt.Print(n, " ")
		}
	}
}