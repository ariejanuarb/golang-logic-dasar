package main

import (
	"fmt"
)

/**
1. membuat looping untuk nilai mulai dari 1 s/d 9, hasil cetak ke kanan 1 baris
2. bentuk 2 dimensi : membuat looping yang sama, namun untuk jumlah 9 baris, dengan lompatan tiap baris (pindah baris)

1. membuat for i
2. membaut for j
3. membuat if (i==j) pada for i
4. membuat else if (i+j) pad a for
*/
func main() {
	//looping untuk membuat baris 1-9
	x := 6
	for j := 1; j <= x; j++ {
		//mulai membuat baris
		n := x
		for i := 1; i <= n; i++ {
			if i == j {
				fmt.Print(j, "-", i, "\t")
			} else if i+j == x+1 {
				fmt.Print(j, "-", i, "\t")
			} else {
				fmt.Print("", "-", "", "\t")
			}
		}
		//membuat baris selesai
		fmt.Println("\n") //pindah baris
	}
}
