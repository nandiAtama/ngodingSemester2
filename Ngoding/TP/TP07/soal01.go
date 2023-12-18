package main

import "fmt"

func inputBilangan_1303223079(bil *int) {
	fmt.Scanln(bil)
	for *bil < 0 {
		fmt.Scanln(bil)
	}
}

func stop_1303223079(bil int) bool {
	return bil == 0
}

func hitung_1303223079(mean *float64, min, max, n *int) {
	var bil int
	inputBilangan_1303223079(&bil)
	*max = bil
	*min = bil
	for !stop_1303223079(bil) {
		if bil > *max {
			*max = bil
		}
		if bil < *min {
			*min = bil
		}
		*mean += float64(bil)
		*n++
		inputBilangan_1303223079(&bil)
	}
}

func main() {
	var min, max, n int
	var mean float64
	hitung_1303223079(&mean, &min, &max, &n)
	if n == 0 {
		fmt.Println("- - - -")
	} else {
		fmt.Print(mean/float64(n), max, min, n)
	}

}