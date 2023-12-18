package main
import "fmt"
type peserta struct {
	negara string
	emas, perak, perunggu, total int
}
const NMAX int = 20
type tabPeserta [NMAX]peserta

func main() {
	var data tabPeserta
	var jumlahNegara int
	var nama string

	jumlahNegara = 11
	bacaData(&data, jumlahNegara)
	// fmt.Println(data)
	fmt.Scan(&nama)
	fmt.Println(cariNegara(data, jumlahNegara, nama))

}

func bacaData(A *tabPeserta, n int) {
	var i int

	for i=0; i<n; i++ {
		fmt.Scanln(&A[i].negara, &A[i].emas, &A[i].perak, &A[i].perunggu, &A[1].total)
	}
}

func cariNegara(A tabPeserta, n int, x string) int {
	var i, idx int
	idx = -2
	i = 1
	for idx == -2 && i < n {
		if A[i].negara == x {
			idx = i
		}
		i++
	}
	return idx + 1
}