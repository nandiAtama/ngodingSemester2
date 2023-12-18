package main

import "fmt"

const NMAX int = 100

type himpunan struct {
	info     [NMAX]string
	nElement int
}

func main() {
	var A, B, C, D himpunan
	createSet_1303223079(&A)
	createSet_1303223079(&B)
	intersection_1303223079(A, B, &C)
	union_1303223079(A, B, &D)
	printSet_1303223079(C)
	printSet_1303223079(D)
}

func createSet_1303223079(set *himpunan) {
	var kata string
	fmt.Scan(&kata)
	for isMember_1303223079(*set, kata) {
		set.info[set.nElement] = kata
		set.nElement++
		fmt.Scan(&kata)
	}
}
func printSet_1303223079(set himpunan) {
	var i int
	for i < set.nElement {
		fmt.Print(set.info[i], " ")
		i++
	}
	fmt.Println()
}
func isMember_1303223079(set himpunan, s string) bool {
	var i int
	for i = 0; i < set.nElement; i++ {
		if s == set.info[i] {
			return false
		}
	}
	return true
}
func intersection_1303223079(set1, set2 himpunan, set3 *himpunan) {
	var i int
	for i < set1.nElement {
		if !isMember_1303223079(set2, set1.info[i]) {
			set3.info[set3.nElement] = set1.info[i]
			set3.nElement++
		}
		i++
	}
}
func union_1303223079(set1, set2 himpunan, set3 *himpunan) {
	var i int
	*set3 = set1
	for i < set2.nElement {
		if isMember_1303223079(*set3, set2.info[i]) {
			set3.info[set3.nElement] = set2.info[i]
			set3.nElement++
		}
		i++
	}
}