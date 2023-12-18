package main

import "fmt"

const NMAX int = 100

type seaGame struct {
	negara                             string
	rank, emas, perak, perunggu, total int
}
type arr [NMAX - 1]seaGame

func main() {
	var tab arr
	var n, m, total, i int
	var name string

	input(&tab, &n)
	fmt.Scan(&name)
	m = binerySearc(tab, n, name)

	if m != -1 {
		fmt.Println("Ada")
	} else {
		fmt.Println("Tidak Ada")
	}

	fmt.Println("Negara", tab[m].negara, "berada diperingkat", tab[m].rank)
	fmt.Println("Negara", tab[m].negara, "mempunyai", tab[m].emas, "medali emas")

	for i = 0; i < n; i++ {
		total += tab[i].total
	}
	fmt.Println(total)
	fmt.Println("Sea games dijuarai oleh negara", tab[max(tab, n)].negara)
}

func input(A *arr, n *int) {
	var input seaGame
	fmt.Scan(&input.rank, &input.negara, &input.emas, &input.perak, &input.perunggu, &input.total)

	for input.negara != "stop" {
		A[*n] = input
		*n++
		fmt.Scan(&input.rank, &input.negara, &input.emas, &input.perak, &input.perunggu, &input.total)
	}
}

func seqSearch(A arr, n int, x string) int {
	var ketemu int = -1
	var i int = 0

	for i < n && ketemu == -1 {
		if A[i].negara == x {
			ketemu = i
		}
		i++
	}
	return ketemu
}

func binerySearc(A arr, n int, x string) int {
	var left, right, mid int
	var idx int
	idx = -1
	left = 0
	right = n - 1
	mid = (left + right) / 2

	for left <= right && idx == -1 {
		if A[mid].negara == x {
			idx = mid
		} else if A[mid].negara > x {
			right = mid - 1
		} else if A[mid].negara < x {
			left = mid + 1
		}
		mid = (left + right) / 2
	}
	return idx
}

func max(A arr, n int) int {
	var idxMax int = 0

	for i := 1; i < n; i++ {
		if A[i].total > A[idxMax].total {
			idxMax = i
		}
	}
	return idxMax
}
