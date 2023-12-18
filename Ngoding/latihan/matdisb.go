package main
import "fmt"

func mencariFaktorial(n int, faktorial *int) {
	var i int
	*faktorial = 1
	for i=1; i <= n; i++ {
		*faktorial = *faktorial * i
	}
}

func permutation(n, r int) int {
	var hasil int
	var f1, f2 int
	mencariFaktorial(n, &f1)
	mencariFaktorial(r, &f2)
	hasil = f1/f2 
	return hasil
}

func combination(n, r int) int {
	var hasil int
	var f1, f2, f3 int

	mencariFaktorial(n, &f1)
	mencariFaktorial(r, &f2)
	mencariFaktorial(n-r, &f3)

	hasil = f1 / (f2*f3)
	return hasil
}

func main() {
	var a, b, c, d int
	var P1, P2, C1, C2 int
	fmt.Scanln(&a, &b, &c, &d)

	P1 = permutation(a, a-c)
	C1 = combination(a, c)
	P2 = permutation(b, b-d)
	C2 = combination(b, d)

	fmt.Println(P1, C1)
	fmt.Println(P2, C2)
}