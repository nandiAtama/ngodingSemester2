package main

import "fmt"

func mencariFaktorial(n int) int {
    var result int = 1
    for i := 1; i <= n; i++ {
        result = result * i
    }
    return result
}

func permutation(n, r int) int {
    return mencariFaktorial(n) / mencariFaktorial(n-r)
}

func combination(n, r in	t) int {
    return mencariFaktorial(n) / (mencariFaktorial(r) * mencariFaktorial(n-r))
}

func main() {
    var n1, n2, r1, r2 int
    fmt.Scan(&n1, &n2, &r1, &r2)
    fmt.Println("P(", n1, ",", r1, ") =", permutation(n1, r1))
    fmt.Println("C(", n1, ",", r1, ") =", combination(n1, r1))
    fmt.Println("P(", n2, ",", r2, ") =", permutation(n2, r2))
    fmt.Println("C(", n2, ",", r2, ") =", combination(n2, r2))

}