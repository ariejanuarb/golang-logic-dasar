package main

import "fmt"

/**
1. membuat looping untuk nilai mulai dari 1 s/d 9, hasil cetak ke kanan 1 baris
2. bentuk 2 dimensi : membuat looping yang sama, namun untuk jumlah 9 baris, dengan lompatan tiap baris (pindah baris)

1. membuat for i
2. membaut for j
3. membuat if (i==j) pada for i
4. membuat else if (i+j) pada for


Soal 1
n = 9
	0	1	2	3	4	5	6	7	8		Identifikasi masalah :
0	1								9		1. Array 1 dimensi membutuhkan 1 for, array 2 dimensi butuh 2 nested for
1		2						8			2. membuat baris angka 1 s/d 9 dengan for loop
2			3				7
3				4		6
4					5
5				4		6
6			3				7
7		2						8
8	1								9

*/

func main() {
	n := 9
	for i := 0; i <= n; i++ {
		fmt.Print(i, "\t")
	}
}

//func main() {
//	//looping untuk membuat baris 1-9
//	x := 9
//	for j := 1; j <= x; j++ {
//		//mulai membuat baris
//		n := x
//		for i := 1; i <= n; i++ {
//			if i == j {
//				fmt.Print(j, "-", i, "\t")
//			} else if i+j == x+1 {
//				fmt.Print(j, "-", i, "\t")
//			} else {
//				fmt.Print("", "-", "", "\t")
//			}
//		}
//		//membuat baris selesai
//		fmt.Println("\n") //pindah baris
//	}
//}
