package main
import "fmt"

var c int
func terurut(a, b *int) {
	c = *a
	*a = *b
	*b = c
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	for x != y {
		if x > y {
			terurut(&x, &y)
		}
		fmt.Println(x, y)
		fmt.Scan(&x, &y)
	}
}