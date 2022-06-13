package logic03

import "fmt"

func Logic03Soal01(n int) { // segitiga kiri = (logic 07 + logic 10)  + logic 05
	a := 3
	nt := n / 2
	for i := 0; i < n; i++ { // loop baris:start
		for j := 0; j < n; j++ { // loop kolom:start
			if i >= j && i+j <= n-1 { // logic 8 reversed && logic 10
				fmt.Print(a, "\t")
			} else if i <= j && i+j >= n-1 { // logic 8 && logic 10 reversed
				fmt.Print(a, "\t")
			} else {
				fmt.Print(" ", "\t")
			}
		}
		fmt.Println("\n")
		if i < nt {
			a += 3
		} else {
			a -= 3
		}
	}
}

func Logic03Soal02(n int) { // sgitiga atas = (logic 08 + logic 10) + logic 05
	a := 3
	nt := n / 2
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i <= j && i+j <= n-1 {
				fmt.Print(a, "\t")
			} else if i >= j && i+j >= n-1 {
				fmt.Print(a, "\t")
			} else {
				fmt.Print(" ", "\t")
			}
		}
		fmt.Println("\n")
		if i < nt {
			a += 3
		} else {
			a -= 3
		}
	}
}
