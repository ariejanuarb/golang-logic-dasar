package logic03

import array2 "github.com/ariejanuarb/golang-logic-dasar/array"

func Logic03Soal03(n int) {
	// 1. create array
	array := array2.NewNumberArray(n, n) // 1. create array

	// 2. isi array
	// loop baris : start
	angka := 3 // int
	for i := 0; i < len(array); i++ {
		nt := n / 2
		// loop kolom : start
		for j := 0; j < len(array[i]); j++ {
			// isi array
			if i < j && j <= nt {
				array[i][j] = int32(angka)
			} else if i > j && j >= nt {
				array[i][j] = int32(angka)
			} else if i+j < n-1 && i >= nt {
				array[i][j] = int32(angka)
			} else if i+j > n-1 && i <= nt {
				array[i][j] = int32(angka)
			}
		}
		// jika baris kurang dari nilai tengah
		if i < nt {
			angka += 3
		} else {
			angka -= 3
		}
	}
	// 3. print array
	array2.PrintNumberArrayZeroEmpty(array)
}
