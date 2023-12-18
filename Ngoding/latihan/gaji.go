package main
import "fmt"

func main() {
	var tj, gaji int
	var jp byte
	fmt.Scanf("%d", &tj)
	fmt.Scanf("%c", &jp)
	gaji = 0
	hitungGajiMingguan(tj, jp, &gaji)
	fmt.Println(gaji)
}

func hitungGajiMingguan(tj int, tp byte, gaji *int) {
    tj-=48
    if tp == 'A'{
        *gaji = 1200000
        if tj >0{
            *gaji+= tj*20000
        }
    }else if tp == 'B'{
        *gaji = 1000000
        if tj >0{
            *gaji+= tj*15000
        }
    }
}