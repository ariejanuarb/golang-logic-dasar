package main

import "fmt"

func main() {
	//Logic02Soal0100(9)
	//Logic02Soal0101(9)
	//Logic02Soal0102(9)
	//Logic02Soal0103(9)

	Logic02Soal0199(9)
}

/*
Soal 01
n = 9
	0	1	2	3	4	5	6	7	8
0	3	3	3	3	3	3	3	3	3
1	6	6	6	6	6	6	6	6	6
2	9	9	9	9	9	9	9	9	9
3	12	12	12	12	12	12	12	12	12
4	15	15	15	15	15	15	15	15	15
5	18	18	18	18	18	18	18	18	18
6	21	21	21	21	21	21	21	21	21
7	24	24	24	24	24	24	24	24	24
8	27	27	27	27	27	27	27	27	27
*/

func Logic02Soal0100(n int) {
	a := 3                   // bilangan awal
	for i := 0; i < n; i++ { // loop baris:start
		for j := 0; j < n; j++ { //loop kolom:start
			fmt.Print(a, "\t")
		} // loop kolom:end
		fmt.Println("\n") // baris baru
		a += 3            // tambah nilai a di setiap baris baru
	} // loop baris:end
}

func Logic02Soal0101(n int) {
	b := 2                   // bilangan awal
	for i := 0; i < n; i++ { // loop baris :start
		for j := 0; j < n; j++ { // loop kolom:start
			fmt.Print(b, "\t") // print b sebanyak nx
		} // loop kolom :end ketika j sudah < n
		fmt.Println("\n")
		b += 3
	}
}

func Logic02Soal0102(n int) {
	c := 4
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(c, "\t")
		}
		fmt.Println("\n")
		c += 2
	}
}

func Logic02Soal0103(n int) {
	d := 5
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(d, "\t")
		}
		fmt.Println("\n")
		d += 3
	}
}

func Logic02Soal0199(n int) {
	for i := 0; i < n; i++ {

	}
}