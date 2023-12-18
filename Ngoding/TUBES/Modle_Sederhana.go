package main

import "fmt"

// USER
type Guru struct {
	nama string
	kode string
	soal Soal
}
type Murid struct {
	nama  string
	nim   int
	nilai int
}
type arrMurid [10]Murid

// SOAL
type Soal struct {
	tugas bankSoal
	quiz  bankSoal
	forum bankSoal
}
type Tugas struct {
	soal, jawaban string
}
type bankSoal [10]Tugas

// ========================================
func inputGuru(T *Guru) {
	var input Guru

	fmt.Print("Masukan nama anda: ")
	fmt.Scan(&input.nama)
	fmt.Print("Masukan kode anda: ")
	fmt.Scan(&input.kode)
	*T = input
}

func inputMurid(T *arrMurid, m *int) {
	var input Murid

	fmt.Print("Masukan nama anda: ")
	fmt.Scan(&input.nama)
	fmt.Print("Masukan NIM anda: ")
	fmt.Scan(&input.nim)
	T[*m] = input
	*m++
}

// ========================================
func tambahKonten(T *Guru, flag string, idxTugas, idxQuiz, idxForum *int) {
	var isi Tugas

	fmt.Println("")
	fmt.Println("Masukan soal & jawaban")
	fmt.Print("Masukan soal: ")
	fmt.Scan(&isi.soal)
	fmt.Print("Masukan jawaban:  ")
	fmt.Scan(&isi.jawaban)
	if flag == "1" || flag == "tugas" {
		T.soal.tugas[*idxTugas] = isi
		fmt.Println("")
		fmt.Println("Isi tugas")
		fmt.Println(T.soal.tugas)
		fmt.Println("Berhasil menambahkan tugas")
		*idxTugas++
	} else if flag == "2" || flag == "quiz" {
		T.soal.quiz[*idxQuiz] = isi
		fmt.Println("")
		fmt.Println("Isi quiz")
		fmt.Println(T.soal.quiz)
		fmt.Println("Berhasil menambahkan quiz")
		*idxQuiz++
	} else if flag == "3" || flag == "forum" {
		T.soal.forum[*idxForum] = isi
		fmt.Println("")
		fmt.Println("Isi forum")
		fmt.Println(T.soal.forum)
		fmt.Println("Berhasil menambahkan forum")
		*idxForum++
	}
}

func editKonten(T *Guru, n int, flag string) {
	var isi Tugas
	var idx int

	fmt.Print("Edit tugas ke: ")
	fmt.Scan(&idx)
	fmt.Print("Masukan soal: ")
	fmt.Scan(&isi.soal)
	fmt.Print("Masukan jawaban: ")
	fmt.Scan(&isi.jawaban)
	if flag == "1" || flag == "tugas" {
		T.soal.tugas[idx] = isi
		fmt.Println("")
		fmt.Println("Isi tugas")
		fmt.Println(T.soal.tugas)
		fmt.Println("Berhasil mengedit tugas")
	} else if flag == "2" || flag == "quiz" {
		T.soal.quiz[idx] = isi
		fmt.Println("")
		fmt.Println("Isi quiz")
		fmt.Println(T.soal.quiz)
		fmt.Println("Berhasil mengedit quiz")
	} else if flag == "3" || flag == "forum" {
		T.soal.forum[idx] = isi
		fmt.Println("")
		fmt.Println("Isi forum")
		fmt.Println(T.soal.forum)
		fmt.Println("Berhasil mengedit forum")
	}
}

func hapusKonten(T *Guru, n *int, flag string) {
	var idx int
	var temp int

	fmt.Print("Hapus konten ke: ")
	fmt.Scan(&idx)
	if flag == "1" || flag == "tugas" {
		T.soal.tugas[idx] = T.soal.tugas[idx+1]
		T.soal.tugas[idx+1] = T.soal.tugas[idx+2]
		T.soal.tugas[idx+2] = T.soal.tugas[idx+3]
		T.soal.tugas[idx+3] = T.soal.tugas[idx+4]
		T.soal.tugas[idx+4] = T.soal.tugas[temp]
		fmt.Println("")
		fmt.Println("Isi tugas")
		fmt.Println(T.soal.tugas)
		fmt.Println("Berhasil menghapus tugas")
		*n--
	} else if flag == "2" || flag == "quiz" {
		T.soal.quiz[idx] = T.soal.quiz[idx+1]
		T.soal.quiz[idx+1] = T.soal.quiz[idx+2]
		T.soal.quiz[idx+2] = T.soal.quiz[idx+3]
		T.soal.quiz[idx+3] = T.soal.quiz[idx+4]
		T.soal.quiz[idx+4] = T.soal.quiz[temp]
		fmt.Println("")
		fmt.Println("Isi quiz")
		fmt.Println(T.soal.quiz)
		fmt.Println("Berhasil menghapus quiz")
		*n--
	} else if flag == "3" || flag == "forum" {
		T.soal.forum[idx] = T.soal.forum[idx+1]
		T.soal.forum[idx+1] = T.soal.forum[idx+2]
		T.soal.forum[idx+2] = T.soal.forum[idx+3]
		T.soal.forum[idx+3] = T.soal.forum[idx+4]
		T.soal.forum[idx+4] = T.soal.forum[temp]
		fmt.Println("")
		fmt.Println("Isi forum")
		fmt.Println(T.soal.forum)
		fmt.Println("Berhasil menghapus forum")
		*n--
	}
}

// ========================================
func sortNim(S *arrMurid, m int) {
	var pass, i, idx int
	var temp Murid
	pass = 1
	for pass < m {
		idx = pass - 1
		i = pass
		for i < m {
			if S[idx].nim > S[i].nim {
				idx = i
			}
			i++
		}
		temp = S[pass-1]
		S[pass-1] = S[idx]
		S[idx] = temp
		pass++
	}
}

func seqSearch(T Guru, n int, kode string) int {

	for i := 0; i < n; i++ {
		if T.kode == kode {
			return i
		}
	}
	return -1
}

func binarysearch(S arrMurid, m int, NIM int) int {
	var left, right, mid, found int
	sortNim(&S, m)
	fmt.Println(S)
	found = -1
	left = 0
	right = m - 1
	for left <= right && found == -1 {
		mid = (left + right) / 2
		if NIM < S[mid].nim {
			right = mid - 1
		} else if NIM > S[mid].nim {
			left = mid + 1
		} else {
			found = mid
		}
	}
	return found
}

func jawabTugas(T Guru, S *arrMurid, m, idxTugas, idxMurid int) {
	var inpJawaban string

	fmt.Print("Masukan jawaban anda: ")
	fmt.Scan(&inpJawaban)
	if inpJawaban == T.soal.tugas[idxTugas].jawaban {
		S[idxMurid].nilai += 100
	}
	fmt.Println(S[idxMurid].nilai)
}

func jawabQuiz(T Guru, S *arrMurid, m, idxQuiz, idxMurid int) {
	var inpJawaban string

	fmt.Print("Masukan jawaban anda: ")
	fmt.Scan(&inpJawaban)
	if inpJawaban == T.soal.quiz[idxQuiz].jawaban {
		S[idxMurid].nilai += 100
	}
	fmt.Println(S[idxMurid].nilai)

}

func jawabForum(T Guru, S *arrMurid, m, idxForum, idxMurid int) {
	var inpJawaban string

	fmt.Print("Jawab forum: ")
	fmt.Scan(&inpJawaban)
	if inpJawaban == T.soal.forum[idxForum].jawaban {
		S[idxMurid].nilai += 100
	} else if inpJawaban != T.soal.forum[idxForum].jawaban {
		S[idxMurid].nilai += 100
	}
	fmt.Println(S[idxMurid].nilai)
}

func selectionSortDesc(S *arrMurid, m int) {
	var pass, i, idx int
	var temp Murid
	pass = 1
	for pass <= m-1 {
		idx = pass - 1
		i = pass
		for i <= m-1 {
			if S[idx].nilai < S[i].nilai {
				idx = i
			}
			i++
		}
		temp = S[pass-1]
		S[pass-1] = S[idx]
		S[idx] = temp
		pass++
	}
}

func insertionSortAsc(S *arrMurid, m int) {
	var pass, i int
	var temp Murid
	pass = 1
	for pass <= m-1 {
		i = pass
		temp = S[pass]
		for i > 0 && temp.nilai < S[i-1].nilai {
			S[i] = S[i-1]
			i--
		}
		S[i] = temp
		pass++

	}
}

func printNilaiDescending(S arrMurid, m, idxMurid int) {
	selectionSortDesc(&S, m)
	for i := 0; i < m; i++ {
		fmt.Println(S[i].nama, S[i].nilai)
	}
}

func printNilaiAScending(S arrMurid, m, idxMurid int) {
	insertionSortAsc(&S, m)
	for i := 0; i < m; i++ {
		fmt.Println(S[i].nama, S[i].nilai)
	}
}

func main() {
	var T Guru
	var S arrMurid
	var idxTugas, idxQuiz, idxForum, n, m int
	idxTugas = 1
	idxQuiz = 1
	idxForum = 1
	var idxMurid int
	var pilihTugas int
	var NIM int
	var reg string
	var inDashboard string
	var flag string
	var num string
	var passGuru string
	var password string = "123123"

	fmt.Println("= = = = = = = = = = = = = = = = = = = = =")
	fmt.Println("=                                       =")
	fmt.Println("=               e-Learning              =")
	fmt.Println("=                                       =")
	fmt.Println("= = = = = = = = = = = = = = = = = = = = =")
	for num != "3" {
		fmt.Print("\n=========================================\n")
		fmt.Print("\nMasuk sebagai\n1. Guru\n2. Murid\n3. Keluar\nPilih: ")
		fmt.Scan(&num)
		if num == "1" {
			fmt.Print("\n=========================================\n")
			fmt.Print("\nMasukan password guru: ")
			fmt.Scan(&passGuru)
			if passGuru == password {
				fmt.Print("\n=========================================\n")
				fmt.Print("\nDASHBOARD\n1. Tambah Konten\n2. Edit Konten\n3. Hapus Konten\n4. Print Nilai Murid\n5. Kembali\n")
				fmt.Print("Pilih: ")
				fmt.Scan(&inDashboard)
				for inDashboard != "5" {

					if inDashboard == "1" {
						fmt.Print("\n=========================================\n")
						fmt.Print("\nMENAMBAHKAN KONTEN\n1. Tambahkan Tugas\n2. Tambahkan Quiz\n3. Tambahkan Forum\nPilih: ")
						fmt.Scan(&flag)
						fmt.Print("\n=========================================\n")
						if flag == "1" || flag == "tugas" {
							tambahKonten(&T, flag, &idxTugas, &idxQuiz, &idxForum)
						} else if flag == "2" || flag == "quiz" {
							tambahKonten(&T, flag, &idxTugas, &idxQuiz, &idxForum)
						} else if flag == "3" || flag == "forum" {
							tambahKonten(&T, flag, &idxTugas, &idxQuiz, &idxForum)
						}

					} else if inDashboard == "2" {
						fmt.Print("\n=========================================\n")
						fmt.Print("\nMENGEDIT KONTEN\n1. Edit Tugas\n2. Edit Quiz\n3. Edit Forum\nPilih: ")
						fmt.Scan(&flag)
						fmt.Print("\n=========================================\n")
						if flag == "1" || flag == "tugas" {
							editKonten(&T, n, flag)
						} else if flag == "2" || flag == "quiz" {
							editKonten(&T, n, flag)
						} else if flag == "3" || flag == "forum" {
							editKonten(&T, n, flag)
						}

					} else if inDashboard == "3" {
						fmt.Print("\n=========================================\n")
						fmt.Print("\nMENGHAPUS KONTEN\n1. Hapus Tugas\n2. Hapus Quiz\n3. Hapus Forum\nPilih: ")
						fmt.Scan(&flag)
						fmt.Print("\n=========================================\n")
						if flag == "1" || flag == "tugas" {
							hapusKonten(&T, &idxTugas, flag)
						} else if flag == "2" || flag == "quiz" {
							hapusKonten(&T, &idxQuiz, flag)
						} else if flag == "3" || flag == "forum" {
							hapusKonten(&T, &idxForum, flag)
						}

					} else if inDashboard == "4" {
						fmt.Print("\n=========================================\n")
						fmt.Println("\nTampilkan data nilai murid berdasarkan")
						fmt.Print("\n1. Nilai terbesar\n2. Nilai terkecil\n")
						fmt.Print("Pilih: ")
						fmt.Scan(&flag)
						fmt.Println()
						if flag == "1" {
							fmt.Println("Data terurut dari nilai terbesar")
							printNilaiDescending(S, m, idxMurid)
						} else if flag == "2" {
							fmt.Println("Data terurut dari nilai terkecil")
							printNilaiAScending(S, m, idxMurid)
						}
					}
					fmt.Print("\n=========================================\n")
					fmt.Print("\nDASHBOARD\n1. Tambah Konten\n2. Edit Konten\n3. Hapus Konten\n4. Print Nilai Murid\n5. Kembali\n")
					fmt.Print("Pilih: ")
					fmt.Scan(&inDashboard)
				}
			}
		} else if num == "2" {
			fmt.Print("\n=========================================\n")
			fmt.Println(S)
			fmt.Print("\nRegistrasi jika belum mempunyai akun\nLogin jika sudah registrasi\n")
			fmt.Print("\n1. Registrasi sebagai murid\n2. Login sebagai murid\n3. Kembali\npilih: ")
			fmt.Scan(&reg)
			for reg != "3" {
				if reg == "1" {
					fmt.Print("\n=========================================\n")
					fmt.Println("\nMasuk sebagai murid")
					inputMurid(&S, &m)
					sortNim(&S, m)
				} else if reg == "2" {
					fmt.Print("\n=========================================\n")
					fmt.Print("\nMasukan NIM anda: ")
					fmt.Scan(&NIM)
					idxMurid = binarysearch(S, m, NIM)
					fmt.Println(idxMurid)
					if idxMurid == -1 {
						fmt.Println("NIM murid tidak ditemukan")
					} else {
						fmt.Print("\n=========================================\n")
						fmt.Print("\nDASHBOARD\n1. Tugas\n2. Quiz\n3. Forum\n4. Kembali\npilih: ")
						fmt.Scan(&inDashboard)
						for inDashboard != "4" {
							for inDashboard != "4" {
								if inDashboard == "1" {
									fmt.Print("\n=========================================\n")
									fmt.Print("\nPilih tugas ke: ")
									fmt.Scan(&pilihTugas)
									fmt.Println(T.soal.tugas[pilihTugas].soal)
									jawabTugas(T, &S, idxTugas, pilihTugas, idxMurid)
								} else if inDashboard == "2" {
									fmt.Print("\n=========================================\n")
									fmt.Print("\nPilih quiz ke: ")
									fmt.Scan(&pilihTugas)
									fmt.Println(T.soal.quiz[pilihTugas].soal)
									jawabQuiz(T, &S, idxTugas, pilihTugas, idxMurid)
								} else if inDashboard == "3" {
									fmt.Print("\n=========================================\n")
									fmt.Print("\nPilih forum ke: ")
									fmt.Scan(&pilihTugas)
									fmt.Println(T.soal.forum[pilihTugas].soal)
									jawabForum(T, &S, idxTugas, pilihTugas, idxMurid)
								}
								fmt.Print("\n=========================================\n")
								fmt.Print("\nDASHBOARD\n1. Tugas\n2. Quiz\n3. Forum\n4. Kembali\npilih: ")
								fmt.Scan(&inDashboard)
							}
						}
					}
				}
				fmt.Print("\n=========================================\n")
				fmt.Println(S)
				fmt.Print("\nRegistrasi jika belum mempunyai akun\nLogin jika sudah registrasi\n")
				fmt.Print("\n1. Registrasi sebagai murid\n2. Login sebagai murid\n3. Kembali\npilih: ")
				fmt.Scan(&reg)
			}
		}
	}
}
