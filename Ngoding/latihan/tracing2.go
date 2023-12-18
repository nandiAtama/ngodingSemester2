package main
import "fmt"

func main() {
	var b int = stone(lily(4))
	fmt.Println(b)
}

func lily(x int) int {
	return 2*x + 5
}

func stone(x int) int {
	var i, y int

	y=0 
	for i=0; i<= x;i++ {
		y = y + 1
	}
	return y
}