package main
import "fmt"

func faktor(n, i int) {
	if i == 1 {
		fmt.Print(1, " ")
	} else {
		faktor(n, i-1)
		if n % i == 0 {
			fmt.Print(i, " ")
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	faktor(n, n)
}