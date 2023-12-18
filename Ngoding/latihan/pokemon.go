package main

import (
	"fmt"
)

const NMAX int = 100

type Pokemon struct {
	name                         string
	HP, Atk, Def, sp_Atk, sp_Def int
	Atk_speed                    float64
}
type arrPokemon [NMAX]Pokemon

func inp(T *arrPokemon, N int) {
	var inp Pokemon

	for i := 0; i < N; i++ {
		fmt.Scan(&inp.name, &inp.HP, &inp.Atk, &inp.Def, &inp.sp_Atk, &inp.sp_Def, &inp.Atk_speed)
		T[i] = inp
	}
}

func print(T arrPokemon, N int) {

	for i := 0; i < N; i++ {
		fmt.Println(T[i].name, T[i].HP, T[i].Atk, T[i].Def, T[i].sp_Atk, T[i].sp_Def, T[i].Atk_speed)
	}
}

func selectionSort_sp_Atk(T *arrPokemon, N int) {
	var pass, idx, i int
	var temp Pokemon

	pass = 1
	for pass < N {
		idx = pass - 1
		i = pass
		for i < N {
			if T[idx].sp_Atk < T[i].sp_Atk {
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

func insertionSort_sp_Atk(T *arrPokemon, N int) {
	var pass, i int
	var temp Pokemon

	pass = 1
	for pass < N {
		i = pass
		temp = T[pass]
		for i > 0 && temp.sp_Atk < T[i-1].sp_Atk {
			T[i] = T[i-1]
			i--
		}
		T[i] = temp
		pass++
	}

}

func seqSearch(T arrPokemon, N int, X string) {
	var found int = -1
	var i int = 0

	for i < N && found == -1 {
		if T[i].name == X {
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

func binSearch(T arrPokemon, N int, X int) string {
	var found int = -1
	var left, mid, right int
	// var max, min, i int

	left = 0
	right = N - 1
	for left <= right && found == -1 {
		mid = (left + right) / 2
		if X < T[mid].sp_Atk {
			right = mid - 1
		} else if X > T[mid].sp_Atk {
			left = mid + 1
		} else {
			found = mid
		}
	}
	return T[found].name
}

func maxAttack(T arrPokemon, N int) string {
	var idx int = 0

	for i := 0; i < N; i++ {
		if T[i].Atk > T[idx].Atk {
			idx = i
		}
	}
	return T[idx].name
}

func minDefend(T arrPokemon, N int) string {
	var idx int = 0

	for i := 0; i < N; i++ {
		if T[i].Def < T[idx].Def {
			idx = i
		}
	}
	return T[idx].name
}

func main() {
	var T arrPokemon
	var n int = 6
	var name string
	var sp_atk int

	inp(&T, n)
	fmt.Println("")
	fmt.Println("Sebelum diurutin")
	print(T, n)

	selectionSort_sp_Atk(&T, n)
	fmt.Println("")
	fmt.Println("Sesudah diurutin menggunakan selection sort")
	print(T, n)

	insertionSort_sp_Atk(&T, n)
	fmt.Println("")
	fmt.Println("Sesudah diurutin menggunakan insertion sort")
	print(T, n)

	fmt.Println("")
	fmt.Print("Masukan nama pokemon yang dicari: ")
	fmt.Scan(&name)
	seqSearch(T, n, name)

	fmt.Println("")
	fmt.Println("Pokemon dengan Attak tertinggi adalah: ", maxAttack(T, n))
	fmt.Println("Pokemon dengan Defend terendah adalah: ", minDefend(T, n))

	fmt.Print("Cari sp_Atk: ")
	fmt.Scan(&sp_atk)
	name = binSearch(T, n, sp_atk)
	fmt.Println(name)

}
