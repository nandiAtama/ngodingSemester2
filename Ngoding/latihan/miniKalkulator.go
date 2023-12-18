package main
import "fmt"

func main() {
	
}

func intro() {
	fmt.Println("Welcome to our app")
	fmt.Scanln()

}

func menu() {
	var pilihan int
	for {
		fmt.Print("===================")
		fmt.Print("        MENU       ")
		fmt.Print("===================")
		fmt.Print("1. Hitung Luas")
		fmt.Print("2. Hitung Keliling")
		fmt.Print("3. Exit")
		fmt.Print("===================")
		fmt.Print("Pilihan anda (1/2/3): ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			hapusLayar()
			hitungLuas()
		case 2:
			hapusLayar()
			hitungKeliling()
		}
		hapusLayar()
		if pilihan == 3 {
			break
		}
	}
}

