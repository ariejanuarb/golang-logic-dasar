package logic02

import (
	"fmt"
)

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

what? print given sequence of number (2 dimmensional array with n=9x9)
who? 1 starting number : a = 3
when? 1 pattern : a += 3 (3,6,9,12,15 so on)
where? conditoning every new row (a += 3) every time loop column is done
why? to create 1 dimensional array using 1 nested for loop + if-else method
how? using for loop method with 2 condition (if & else) to create 2 pattern in 1 dimensional array
*/

func Logic02Soal0100(n int) {
	a := 3                   // starting number
	for i := 0; i < n; i++ { // loop baris:start
		for j := 0; j < n; j++ { //loop kolom:start
			fmt.Print(a, "\t") // print a 9x on every column
		} // loop kolom:end
		fmt.Println("\n") // jump into new rows
		a += 3            // add a by 3 on every new rows
	} // loop baris:end
}

func Logic02Soal0101(n int) {
	b := 3                   // bilangan awal
	for i := 0; i < n; i++ { // loop baris :start
		for j := 0; j < n; j++ { // loop kolom:start
			fmt.Print(b, "\t") // print b sebanyak nx
		} // loop kolom :end ketika j sudah < n
		fmt.Println("\n")
		b += 3
	}
}

func Logic02Soal0102(n int) {
	c := 3
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(c, "\t")
		}
		fmt.Println("\n")
		c += 3
	}
}

func Logic02Soal0103(n int) {
	d := 3
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(d, "\t")
		}
		fmt.Println("\n")
		d += 3
	}
}

func Logic02Soal0104(n int) {
	e := 3
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(e, "\t")
		}
		fmt.Println("\n")
		e += 3
	}
}

func Logic02Soal0105(n int) {
	f := 3
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(f, "\t")
		}
		fmt.Println("\n")
		f += 3
	}
}
