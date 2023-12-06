package main

import "fmt"

func main() {
	var orang, dewasa, kecil, sisa int
	fmt.Scanln(&orang)

	dewasa = orang / 5
	sisa = dewasa % 5

	if dewasa <= 3 && orang <= 15 {
		if sisa == 0 {
			fmt.Println("dewasa", dewasa)
		} else {
			fmt.Println("dewasa", dewasa+1)
		}
	} else {
		dewasa = 3
		orang = orang - 15
		kecil = orang / 2
		sisa = orang % 2
		if kecil <= 5 && orang <= 10 {
			if sisa == 0 {
				fmt.Println("Dewasa", dewasa, "kecil", kecil)
			} else {
				fmt.Println("Dewasa", dewasa, "kecil", kecil+1)
			}
		} else {
			kecil = 5
			orang = orang - 10
			fmt.Println("Dewasa", dewasa, "Kecil", kecil, "dan", orang, "tak berangkat")
		}
	}

}
