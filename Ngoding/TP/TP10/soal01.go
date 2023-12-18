package main

import (
	"fmt"
)

const NMAX int = 12345

type prodi struct {
	nama, akreditasi      string
	tahun, aktif, lulusan int
}
type fakultas struct {
	nama     string
	arrProdi [NMAX]prodi
	N        int
}

func main() {
	var fif fakultas
	var j int
	var terbanyak prodi
	var termuda int
	var prestasi string
	fmt.Scanln(&fif.nama)
	for j = 0; j < 10; j++ {
		dataProdi_1303223079(&fif)
	}
	terbanyak = terbanyak_1303223079(fif)
	fmt.Println("Prodi", terbanyak.nama, "memiliki mahasiswa dan lulusan terbanyak yaitu", terbanyak.aktif+terbanyak.lulusan)
	termuda = termuda_1303223079(fif)
	fmt.Println("Prodi termuda adalah", fif.arrProdi[termuda].nama)
	prestasi = prestasi_1303223079(fif)
	fmt.Println("Akreditasi prodi terbanyak di Fakultas", fif.nama, "adalah", prestasi)
	cetakProdi_1303223079(fif, prestasi)

}

func buatFakultas_1303223079(f *fakultas) {
	fmt.Scanln(f.nama)
	f.N = 0
}

func dataProdi_1303223079(f *fakultas) {
	var p prodi
	var indeks int
	fmt.Scanln(&p.nama, &p.akreditasi, &p.tahun, &p.aktif, &p.lulusan)
	indeks = cekProdi_1303223079(p.nama, *f)
	if indeks == -1 {
		f.arrProdi[f.N] = p
		f.N++
	} else {
		f.arrProdi[indeks].aktif += p.aktif
		f.arrProdi[indeks].lulusan += p.lulusan
	}
}

func cekProdi_1303223079(s string, f fakultas) int {
	var j int = 0
	for j < f.N && f.arrProdi[j].nama != s {
		j++
	}
	if j == f.N {
		return -1
	} else {
		return j
	}
}

func terbanyak_1303223079(f fakultas) prodi {
	var banyak prodi = f.arrProdi[0]
	for i := 1; i < f.N; i++ {
		if banyak.aktif+banyak.lulusan < f.arrProdi[i].aktif+f.arrProdi[i].lulusan {
			banyak = f.arrProdi[i]
		}
	}
	return banyak
}

func termuda_1303223079(f fakultas) int {
	var muda int = 0
	for i := 1; i < f.N; i++ {
		if f.arrProdi[muda].tahun <= f.arrProdi[i].tahun {
			muda = i
		}
	}
	return muda
}

func prestasi_1303223079(f fakultas) string {
	var unggul, baik, cukup int
	for i := 0; i < f.N; i++ {
		if f.arrProdi[i].akreditasi == "unggul" {
			unggul++
		} else if f.arrProdi[i].akreditasi == "baik" {
			baik++
		} else if f.arrProdi[i].akreditasi == "cukup" {
			cukup++
		}
	}
	if unggul > baik && unggul > cukup {
		return "unggul"
	} else if baik > unggul && baik > cukup {
		return "baik"
	} else {
		return "cukup"
	}
}

func cetakProdi_1303223079(f fakultas, x string) {
	for i := 0; i < f.N; i++ {
		if f.arrProdi[i].akreditasi == x {
			fmt.Print(f.arrProdi[i].nama, " ")
		}
	}
}