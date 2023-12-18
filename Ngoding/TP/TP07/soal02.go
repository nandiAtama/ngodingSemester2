package main
import "fmt"

func average_1303223079(bil, i int, hasil *float64) {
	if bil == 0 {
		*hasil = *hasil / float64(i)
	}else{
		i++
		*hasil += float64(bil%10)
		fmt.Println(*hasil)
		average_1303223079(bil/10,i,hasil)
	}
	
}

func main() {
	var inp, n int
	var hasil float64
	fmt.Scan(&inp)
	average_1303223079(inp, n, &hasil)
	fmt.Println(hasil)
}