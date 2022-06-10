package logic02

import array2 "github.com/ariejanuarb/golang-logic-dasar/array"

func Logic02Soal0111(n int) {
	// 1. create array
	array := array2.NewNumberArray(n, n)

	// 2. isi array
	angka := 3                        // angka pertama
	for i := 0; i < len(array); i++ { // loop baris : start
		// loop kolom : start
		for j := 0; j < len(array[i]); j++ {
			// isi array
			array[i][j] = int32(angka)
		}
	}
	// 3. print array
	array2.PrintNumberArray(array)
}
