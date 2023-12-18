package main
import ( 
	"fmt"
	"math"
)

type titik struct {
	x, y float64
}

func main() {
	var p1, p2 titik

	fmt.Scan(&p1.x, &p1.y, &p2.x, &p2.y)
	fmt.Printf("%.2f",jarak(p1, p2))
}

func jarak(p1, p2 titik) float64 {
	return akar(((p1.x-p2.x)*(p1.x-p2.x)) + ((p1.y-p2.y)*(p1.y-p2.y)))
}

func akar(x float64) float64 {
	return math.Sqrt(x)
}