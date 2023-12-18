package main

import "fmt"

type motor struct {
	merek, warna string
	cc           int
}

const NMAX int = 10

type TabMotor [NMAX]motor

func main() {
	var ride TabMotor
	var n int

	fmt.Scan(&n)
	bacaData(&ride, n)
	fmt.Println(cc_terbesar(ride, n))
	fmt.Println(merek_cc_terbesar(ride, n))
	fmt.Println(warna_cc_terbesar(ride, n))
}

func bacaData(r *TabMotor, n int) {
	var i int = 0
	for i < n {
		fmt.Scan(&r[i].merek, &r[i].warna, &r[i].cc)
		i++
	}
}

func idx_cc_terbesar(r TabMotor, n int) int {
	var i, idx_max int
	for i = 1; i < n; i++ {
		if r[i].cc > r[idx_max].cc {
			idx_max = i
		}
	}
	return idx_max
}

// func cc_terbesar(r TabMotor, n int) int {
// 	var i, cc_max int
// 	cc_max = r[0].cc
// 	for i = 1; i < n; i++ {
// 		if r[i].cc > cc_max {
// 			cc_max = r[i].cc
// 		}
// 	}
// 	return cc_max
// }
func cc_terbesar(r TabMotor, n int) int {
	var idx_max int
	idx_max = idx_cc_terbesar(r, n)
	return r[idx_max].cc
}

func merek_cc_terbesar(r TabMotor, n int) string {
	var idx_max int
	idx_max = idx_cc_terbesar(r, n)
	return r[idx_max].merek
}

func warna_cc_terbesar(r TabMotor, n int) string {
	var idx_max int
	idx_max = idx_cc_terbesar(r, n)
	return r[idx_max].warna
}