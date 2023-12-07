package main

import "fmt"

func main() {
	var member string
	var a, b, c, d, e int
	var cashback, diskon float64

	fmt.Scanln(&member, &a, &b, &c, &d, &e)
	if member == "Yes" {
		if a%2 != 0 && b%2 != 0 && c%2 != 0 && d%2 != 0 && e%2 != 0 {
			diskon = 0.15*(float64(c+d+e)*1.7) + (float64(c+d+e))*1.7
		} else if a%2 == 0 && b%2 == 0 && c%2 == 0 && d%2 == 0 && e%2 == 0 {
			cashback = 0.15*(float64(a+b+c))*3.1 + (float64(a+b+c))*3.1
		} else {
			diskon = 0.15*(float64(c+d+e)*1.7) + (float64(c+d+e))*1.7
			cashback = 0.15*(float64(a+b+c))*3.1 + (float64(a+b+c))*3.1
		}
		if diskon >= 50.0 {
			diskon = 50.0
		}
		if cashback >= 35 {
			cashback = 35
		}
	} else {
		if a%2 != 0 && b%2 != 0 && c%2 != 0 && d%2 != 0 && e%2 != 0 {
			diskon = (float64(c + d + e)) * 1.7
		} else if a%2 == 0 && b%2 == 0 && c%2 == 0 && d%2 == 0 && e%2 == 0 {
			cashback = (float64(a + b + c)) * 3.1
		} else {
			diskon = (float64(c + d + e)) * 1.7
			cashback = (float64(a + b + c)) * 3.1
		}
		if diskon >= 50.0 {
			diskon = 50.0
		}
		if cashback >= 35 {
			cashback = 35
		}
	}
	fmt.Println("Cashback:", cashback, "%", "Diskon:", diskon, "%")
}
