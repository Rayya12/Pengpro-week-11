package main

import "fmt"

func main() {
	var JP, MP int
	var tarif, JL float64
	fmt.Scanln(&JP, &MP, &JL)
	if (JP >= 5 && JP < 9) || (JP == 9 && MP == 0) || (JP >= 16 && JP < 19) || (JP == 19 && MP == 0) {
		if JL > 0.0 && JL <= 10.0 {
			tarif = 5000.0 * JL
			fmt.Println(tarif)
		} else {
			tarif = 4500.0 * JL
			fmt.Println(tarif)
		}
	} else if (JP > 19 && JP < 22) || (JP == 22 && MP == 0) || (JP > 9 && JP < 16) {
		tarif = 4000.0 * JL
		fmt.Println(tarif)
	} else {
		fmt.Println("maaf, kami tidak bisa melayani pesanan anda")
	}
}
