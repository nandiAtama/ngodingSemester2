package main
import "fmt"


type pokemon struct {
	nama string
	CP, HP, IV_Atk, IV_Def, IV_HP int
}
type arrPokemon [2023]pokemon


func inpPokemon_1303223079(T *arrPokemon, N *int) {
	var inp pokemon

	fmt.Scanln(&*N)
	for i := 0; i < *N; i++ {
		fmt.Scan(&inp.nama, &inp.CP, &inp.HP, &inp.IV_Atk, &inp.IV_Def, &inp.IV_HP)
		T[i] = inp
	}

}

func printPokemon_1303223079(T arrPokemon, N int) {

	for i := 0; i < N; i++ {
		fmt.Println(T[i].nama, T[i].CP, T[i].HP, T[i].IV_Atk, T[i].IV_Def, T[i].IV_HP)
	}

}

func urutPokemon_1303223079(T *arrPokemon, N int, flag string) {
	var pass, i int
	var temp pokemon
	pass = 1

	for pass < N {
		i = pass
		temp = T[pass]
		if flag == "CP" {
			for i > 0 && temp.CP > T[i-1].CP {
				T[i] = T[i-1]
				i--
			}
		} else if flag == "HP" {
			for i > 0 && temp.HP > T[i-1].HP {
				T[i] = T[i-1]
				i--
			}
		} else if flag == "IV_Atk" {
			for i > 0 && temp.IV_Atk > T[i-1].IV_Atk {
				T[i] = T[i-1]
				i--
			}
		} else if flag == "IV_Def" {
			for i > 0 && temp.IV_Def > T[i-1].IV_Def {
				T[i] = T[i-1]
				i--
			}
		} else if flag == "IV_HP" {
			for i > 0 && temp.IV_HP > T[i-1].IV_HP {
				T[i] = T[i-1]
				i--
			}
		}	
		T[i] = temp
		pass++
	}
}

func totalIV_1303223079(T arrPokemon, N int, namaX string) float64 {
	var total float64
	var idx int

	for i := 0; i < N; i++ {
		if T[i].nama == namaX {
			idx = i
			total = (float64(T[idx].IV_Atk + T[idx].IV_Def + T[idx].IV_HP) * 100) / 45
			return total
		}	
	}
	return -9999
}

func main() {
	var T arrPokemon
	var N int
	var namaX, flag string

	inpPokemon_1303223079(&T, &N)
	fmt.Scan(&namaX, &flag)
	fmt.Println()
	printPokemon_1303223079(T, N)
	fmt.Println()
	urutPokemon_1303223079(&T, N, flag)
	printPokemon_1303223079(T, N)
	fmt.Println()
	total := totalIV_1303223079(T, N, namaX)
	fmt.Printf("Total IV %s adalah %.3f", namaX, total)
}