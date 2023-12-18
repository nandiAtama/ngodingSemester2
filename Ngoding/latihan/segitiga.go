package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if segitiga(a, b, c) {
		fmt.Println("Segitia")
	} else {
		fmt.Println("Bukan segitiga")
	}
}
func segitiga(a, b, c int) bool {
	var cek bool
	cek = a+b>c && a+c>b && b+c>a && a!=0 && b!=0 && c!=0 
	return cek
}