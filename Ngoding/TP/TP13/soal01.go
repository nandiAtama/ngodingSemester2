package main
import "fmt"

type student struct {
    name string
    math, indo, eng, sains, average float64
}

type arrStudent [2048] student

func main() {
    var murid arrStudent
    var n int

    entryStudent_1303223079(&murid, &n)
    calculateAverage_1303223079(&murid, n)
    rangking_1303223079(&murid, n)
    printTop3_1303223079(murid, n)
    printMax_1303223079(murid, n)
}

func entryStudent_1303223079(T *arrStudent, n *int) {
    var dataNilai student
    fmt.Scan(&dataNilai.name)
    for *n < 2048 && dataNilai.name != "STOP" {
        fmt.Scan(&dataNilai.math, &dataNilai.indo, &dataNilai.eng, &dataNilai.sains)
        T[*n] = dataNilai
        *n++
        fmt.Scan(&dataNilai.name)
    }
}

func calculateAverage_1303223079(T *arrStudent, n int) {
    var i int = 0
    for i < n {
        T[i].average = (T[i].math + T[i].indo + T[i].eng + T[i].sains) / 4
        i++
    }
}

func max_1303223079(T arrStudent, n int, flag string) int {
    var i int = 0
    var idxMax int = 0
    for i < n {
        
        if flag == "math" {
            if T[idxMax].math < T[i].math {
                idxMax = i
            }
        } else if flag == "indo" {
            if T[idxMax].indo < T[i].indo {
                idxMax = i
            }
        } else if flag == "eng" {
            if T[idxMax].eng < T[i].eng {
                idxMax = i
            }
        } else if flag == "sains" {
            if T[idxMax].sains < T[i].sains {
                idxMax = i
            }
        }
        i++
    }
    return idxMax
}

func rangking_1303223079(T *arrStudent, n int) {
    var pass, idx, i int
    var temp student

    for pass=1; pass <= n-1; pass++ {
        idx = pass-1
        for i = pass; i < n; i++ {
            if T[idx].average < T[i].average {
                idx = i
            }
        }
        temp = T[pass-1]
        T[pass-1] = T[idx]
        T[idx] = temp
    } 
}

func printTop3_1303223079(T arrStudent, n int) {
   var i int

   for i=0; i < 3; i++ {
        fmt.Println("Rangking", i+1, ":", T[i].name, " rata-rata ", T[i].average)
   }
}

func printMax_1303223079(T arrStudent, n int) {
    var maxMath, maxIndo, maxEng, maxSains int
    maxMath = max_1303223079(T, n, "math")
    maxIndo = max_1303223079(T, n, "indo")
    maxEng = max_1303223079(T, n, "eng")
    maxSains = max_1303223079(T, n, "sains")
    fmt.Println("Nilai Matematika tertinggi diraih oleh", T[maxMath].name)
    fmt.Println("Nilai Indonesia tertinggi diraih oleh", T[maxIndo].name)
    fmt.Println("Nilai Inggris tertinggi diraih oleh", T[maxEng].name)
    fmt.Println("Nilai Sains tertinggi diraih oleh", T[maxSains].name)
}