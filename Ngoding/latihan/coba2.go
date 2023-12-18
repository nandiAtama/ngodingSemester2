package main

import (
	"fmt"
)

type Pokemon struct {
	name                                string
	HP, Attack, Defense, SP_atk, SP_def int
	Atk_speed                           float64
}

type arrPokemon [20]Pokemon

func input(T *arrPokemon, n int) {

	for i := 0; i < n; i++ {
		fmt.Scan(&T[i].name, &T[i].HP, &T[i].Attack, &T[i].Defense, &T[i].SP_atk, &T[i].SP_def, &T[i].Atk_speed)
	}
}

func print(T arrPokemon, n int) {

	for i := 0; i < n; i++ {
		fmt.Println(T[i].name, T[i].HP, T[i].Attack, T[i].Defense, T[i].SP_atk, T[i].SP_def, T[i].Atk_speed)
	}
}

func selectionSort(T *arrPokemon, n int) {
	var pass, i, idx int
	var temp Pokemon

	pass = 1
	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if T[idx].SP_atk < T[i].SP_atk {
				idx = i
			}
			i++
		}
		temp = T[pass-1]
		T[pass-1] = T[idx]
		T[idx] = temp
		pass++
	}
}

func insertionSot(T *arrPokemon, n int) {
	var pass, i int
	var temp Pokemon

	pass = 1
	for pass < n {
		i = pass
		temp = T[pass]
		for i > 0 && temp.SP_atk < T[i-1].SP_atk {
			T[i] = T[i-1]
			i--
		}
		T[i] = temp
		pass++
	}
}

func max(T arrPokemon, n int) string {
	var idx int

	for i := 0; i < n; i++ {
		if T[i].Attack > T[idx].Attack {
			idx = i
		}
	}
	return T[idx].name
}

func min(T arrPokemon, n int) string {
	var idx int

	for i := 0; i < n; i++ {
		if T[i].Defense < T[idx].Defense {
			idx = i
		}
	}
	return T[idx].name
}

func seqSearch(T arrPokemon, n int, x string) {
	var found = -1
	var i int

	for i < n && found == -1 {
		if T[i].name == x {
			found = i
		}
		i++
	}
	if found != -1 {
		fmt.Println("Pokemon", T[found].name, "ada")
	} else {
		fmt.Println("Pokemon tidak ada")
	}
}

func binSearch(T arrPokemon, n, x int) string {
	var found int = -1
	var left, mid, right int

	left = 0
	right = n - 1
	for left <= right && found == -1 {
		mid = (left + right) / 2
		if x < T[mid].SP_atk {
			right = mid - 1
		} else if x > T[mid].SP_atk {
			left = mid + 1
		} else {
			found = mid
		}
	}
	if found == -1 {
		kiri := x - (found - 1)
		kanan := (found + 1) - x
		if kiri < kanan {
			return T[kiri].name
		} else {
			return T[kanan].name
		}
	}
	return T[found].name
}

func main() {
	var T arrPokemon
	var n int = 6
	var x string
	var sp_Atk int
	input(&T, n)
	fmt.Println("")
	fmt.Println("Sebelum sort")
	print(T, n)

	selectionSort(&T, n)
	fmt.Println("")
	fmt.Println("Sesudah sort")
	print(T, n)

	insertionSot(&T, n)
	fmt.Println("")
	fmt.Println("sesudah sort")
	print(T, n)

	fmt.Println("")
	fmt.Println("Attack tertinggi: ", max(T, n))
	fmt.Println("Defend terendah: ", min(T, n))

	fmt.Println("")
	fmt.Print("Masukan nama pokemon yang dicari: ")
	fmt.Scan(&x)
	seqSearch(T, n, x)

	fmt.Println("")
	fmt.Print("Cari SP_atk: ")
	fmt.Scan(&sp_Atk)
	x = binSearch(T, n, sp_Atk)
	fmt.Println(x)

}

