package main

import "fmt"

func main() {
	var JD, MD, JP, MP int
	var jam, menit int

	fmt.Scanln(&JD, &MD, &JP, &MP)
	jam = JP - JD
	menit = MP - MD
	if jam < 0 {
		jam = JP + 12 - JD
	}
	if menit < 0 {
		jam = jam - 1
		menit = MP + 60 - MD
	}
	fmt.Println(jam, "jam", menit, "menit")

}
