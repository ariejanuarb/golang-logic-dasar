package logic3

import array2 github.com/ariejanuarb/golang-logic/array

func Logic03Soal03(n int) {
	array := array2.NewNumberArray(n, n) // 1. create array

	// 2. isi array
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {

		}
	}

	// 3. print array
	array2.PrintNumberArray(array)
}