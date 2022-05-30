package main

import "fmt"

func main() {
	//Logic02Soal01(9)
	//Logic02Soal02(9)
	//Logic02Soal03(9)
	//Logic02Soal04(9)
	Logic02Soal05(9)
}

func Logic02Soal01(n int) {
	a := 3
	// loop baris
	for i := 0; i < n; i++ {
		//loop kolom
		for j := 0; j < n; j++ {
			fmt.Print(a, "\t")
		} // kolom selesai
		// baris baru
		fmt.Println("\n")
		// tambah nilai a di baris selanjutnya
		a += 3
	}
}

func Logic02Soal02(n int) {
	// loop baris start
	for i := 0; i < n; i++ {
		// loop kolom
		a := 3
		for j := 0; j < n; j++ {
			fmt.Print(a, "\t")
			a += 3
		}
		// pindah baris baru
		fmt.Println("\n")
		// loop baris ends
	}
}

func Logic02Soal03(n int) {
	// loop baris x9 start
	for i := 0; i < n; i++ {
		// loop kolom x9 start
		a := 3 * n
		for j := 0; j < n; j++ {
			fmt.Print(a, "\t")
			a -= 3
		} // loop kolom end
		// pindah baris baru setiap loop kolom selesai
		fmt.Println("\n")
	} // loop baris ends
}

func Logic02Soal04(n int) {
	a := 3 * n
	// loop baris
	for i := 0; i < n; i++ {
		//loop kolom
		for j := 0; j < n; j++ {
			fmt.Print(a, "\t")
		} // kolom selesai
		// baris baru
		fmt.Println("\n")
		// tambah nilai a di baris selanjutnya
		a -= 3
	}
}

func Logic02Soal05(n int) {
	a := 3
	for i := 0; i < n; i++ {
		fmt.Print(a, "\t")
	}
}
