package logic03

import (
	array2 "github.com/ariejanuarb/golang-logic-dasar/array"
)

func Logic03Soal05arr00(n int) {
	// create array
	array := array2.NewNumberArray(n, n)
	// initial angka dari 3
	angka := 3
	mid := n / 2
	// loop baris
	for b := 0; b < n; b++ {
		// loop kolom
		for k := 0; k < n; k++ { // loop kolom : selesai
			// jika baris 0
			if b%4 == 0 {
				array[k][b] = int32(angka) // input dibalik, [kolom][baris]
			} else if b%4 == 1 && k == n-1 {
				array[k][b] = int32(angka)
			} else if b%4 == 2 {
				array[n-1-k][b] = int32(angka)
			} else if b%4 == 3 && k == 0 {
				array[k][b] = int32(angka)
			}
		}
		if b < mid {
			angka += 3
		} else {
			angka -= 3
		}
	}
	// print array
	array2.PrintNumberArrayZeroEmpty(array)
}
