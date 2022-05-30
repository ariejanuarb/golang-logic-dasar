package main

import "fmt"

/**
Soal 2
n = 9
	0	1	2	3	4	5	6	7	8
0	1	3	5	7	9	11	13	15	17
1	1								17
2	1								17
3	1								17
4	1								17
5	1								17
6	1								17
7	1								17
8	1	3	5	7	9	11	13	15	17
*/

/*func main() {
	n := 9
	for j := 1; j <= n; j++ {
		for i := 1; i <= 17; i += 2 {
			if j == 1 || j == 9 {
				fmt.Print(i, "\t")
			} else if i == 1 || i == 17 {
				fmt.Print(i, "\t")
			} else {
				fmt.Print("", "", "", "\t")
			}
		}
		fmt.Println("\n") // cetak baris baru (kebawah) ketika looping ke-i slesai
	}
}*/

func main() {
	n := 9
	for j := 1; j <= n; j++ {
		for i := 1; i <= 17; i += 2 {
			if j == 1 || j == 9 {
				fmt.Print(i, "\t")
		}
		fmt.Println("\n") // cetak baris baru (kebawah) ketika looping ke-i slesai
	}
}
