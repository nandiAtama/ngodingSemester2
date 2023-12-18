package main
import "fmt"

const nMAX int = 2022
type mahasiswa struct {
    nama, nim string
    gpa float64
}
type tabMhs [nMAX]mahasiswa

func sortGPA(mhs *tabMhs) {
    var pass, idx, i int
    var temp mahasiswa

    for pass = 0 < n {
        idx = pass
        for i = pass+1 < n {
            if mhs[idx].gpa > mhs[idx].gpa {
                idx = i
            }
        }
        temp = mhs[pass]
        mhs[pass] = mhs[idx]
        mhs[idx] = temp
    }
}

func main() {
    var i, n int
    var mhs tabMhs
    n = 10

    for i = 0; i <= n-1; i++ {
        fmt.Scanln(mhs[i].nim, mhs[i])
    }
}